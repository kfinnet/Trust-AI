package main

import (
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {

	// CSV 파일을 연다
	passengersFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV 파일로부터 dataframe을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// log_ditterenced_passengers 열에서 시간 및 
	// 승객 데이터를 floats 슬라이스로 읽어온다
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// autocorrelations 자기상관을 도표로 그리기 위해 새로운 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Partial Autocorrelations for log(differenced passengers)"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "PACF"
	p.Y.Min = 15
	p.Y.Max = -1

	w := vg.Points(3)

	// 도표에 대한 지점을 생성한다.
	numLags := 20
	pts := make(plotter.Values, numLags)

	// 시계열에서 여러 이전 값들을 루프를 통해 읽는다.
	for i := 1; i <= numLags; i++ {

		// 자기상관을 계산한다.
		pts[i-1] = pacf(passengers, i)
	}

	// 앞서 계산한 지점들을 도표에 추가한다.
	bars, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(1)

	// 도표를 PNG 파일로 저장한다.
	p.Add(bars)
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "pacf.png"); err != nil {
		log.Fatal(err)
	}
}

// pacf 함수는 주어진 특정 주기 이전의 값에서
// 시계열의 편 자기상관을 계산한다.
func pacf(x []float64, lag int) float64 {

	// github.com/sajari/regression 사용해 모델을 
	// 훈련(학습)시키기 위해 필요한 regression.Regression 값을 생성한다.
	var r regression.Regression
	r.SetObserved("x")

	// 현재 및 중간 이전의 값을 모두 정의한다.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// 데이터 열을 이동시킨다.
	xAdj := x[lag:len(x)]

	// 루프를 통해 회귀분석 모델을 위한
	// 데이터 집합을 생성하는 시계열 데이터를 읽는다.
	for i, xVal := range xAdj {

		// 루프를 통해 독립 변수를 구성하기 위해 필요한
		// 중간 이전의 값을 읽는다.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// 이전 값들에 대한 시계열 데이터를 얻는다.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// 이 지점들을 regression 값에 추가한다.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// 회귀분석 모델을 훈련(학습)시킨다.
	r.Run()

	return r.Coeff(lag)
}
// 그런다음, 이 pacf 함수를 사용해 일부 값에 대한 편 자기상관을 
// 계산할수 있다.
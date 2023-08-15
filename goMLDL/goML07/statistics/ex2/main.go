package main

import (
	"log"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/gonum/stat"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV 파일로부터 dataframe을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// AirPassengers 열에서 시간 및 승객 데이터를 floats 배열로 읽어온다.
	passengers := passengersDF.Col("AirPassengers").Float()

	// 자기 상관을 도표로 그리기 위해 새로운 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Autocorrelations for AirPassengers"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "ACF"
	p.Y.Min = 0
	p.Y.Max = 1

	w := vg.Points(3)

	// 도표에 대한 지점을 생성한다.
	numLags := 20
	pts := make(plotter.Values, numLags)

	// 시계열에서 여러 이전 값들을 루프를 통해 읽는다.
	for i := 1; i <= numLags; i++ {

		// 자기상관을 계산한다.
		pts[i-1] = acf(passengers, i)
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
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "acf.png"); err != nil {
		log.Fatal(err)
	}
}

// acf 함수는 주어진 이전 데이터와의 구간에서
// 시계열에 대한 자기상관(autocorrelation)을 계산한다.
func acf(x []float64, lag int) float64 {

	// 시계열을 이동시킨다.
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator 변수는 누적된 분자의 값을 저장하는데 사용되며,
	// denominator 변수는 누적된 분모의 값을 저장하는 데 사용된다.
	var numerator float64
	var denominator float64

	// 자기상관(autocorrelation)의 각 항에 사용될
	// x 값의 평균을 계산한다.
	xBar := stat.Mean(x, nil)

	// numerator(분자)를 계산한다.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// denominator(분모)를 계산한다.
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}

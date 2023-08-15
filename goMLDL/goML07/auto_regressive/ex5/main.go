package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	// CSV 파일로 부터 dataframe 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// log_differenced_passengers 열에서 시간 및 승객 데이터를
	// floats 슬라이스로 읽어온다.
	passengers := passengersDF.Col("log_differenced_passengers").Float()

	// 1주기, 2주기 이전 관찰 값과 오류에 대한 계수를 계산한다.
	coeffs, intercept := autoregressive(passengers, 2)

	// 표준출력을 통해 AR(2) 모델을 출력한다.
	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f + lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])
}

// autoregressive 함수는 주어진 순서(order) 이전 시점의
// 시계열 데이터에 대한 AR 모델을 계산한다.
func autoregressive(x []float64, lag int) ([]float64, float64) {

	// github.com/sajari/regression 사용해 모델을
	// 훈련(학습)시키지 위해 필요한 regression.Regression 값을 생성한다.
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

		// 이 지점들을 regressin 값에 추가한다.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// 회귀분석 모델을 훈련(적합)시킨다.
	r.Run()

	// coeff는 이전 주기에 대한 계수를 저장한다.
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}

	return coeff, r.Coeff(0)
}
// 이제 로그 변환 및 차분을 거친 시계열에서 이 함수를 호출하면
// 학습을 마친 AR(2) 모델의 계수를 구할 수 있다.
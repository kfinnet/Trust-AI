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
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// dataframe을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// AirPassengers 열에서 시간 및 승객 데이터를 float 배열로 읽어온다.
	passengers := passengersDF.Col("AirPassengers").Float()

	// 루프를 통해 시계열에서 다양한 주기의 이전 값들을 읽는다
	fmt.Println("Partial Autocorrelation편 자기상관:")
	for i := 1; i < 11; i++ {

		// 편 자기상관을 계산한다.
		pac := pacf(passengers, i)
		fmt.Printf("%d period 주기 이전 값: %0.2f\n", i, pac)
	}
}
// 컴파일 실행하면 항공 승객 시계열에 대한 편 자기상관의 값을 구할수 있다

// pacf 함수는 주어진 특정 주기 이전의 값에서
// 시계열의 편 자기상관을 계산한다.
func pacf(x []float64, lag int) float64 {

	// github.com/sajari/regression 을 사용해 모델을
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

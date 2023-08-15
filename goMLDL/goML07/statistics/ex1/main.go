package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"gonum.org/v1/gonum/stat"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다.
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV 파일로 부터 dataframe 을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// AirPassengers 열에서 시간 및 승객 데이터를 floats 배열로 읽어온다.
	passengers := passengersDF.Col("AirPassengers").Float()

	// 시계열에서 여러 이전 값들을 루프를 통해 읽는다.
	fmt.Println("Autocorrelation자기상관:")
	for i := 1; i < 11; i++ {

		// 자기 상관을 계산한다.
		ac := acf(passengers, i)
		fmt.Printf("%d 주기이전 값period: %0.2f\n", i, ac)
	}
}

// acf 함수는 주어진 이전 데이터와의 구간에서
// 시계열에 대한 자기상관(autocorrelation)을 계산한다.
func acf(x []float64, lag int) float64 {

	// 시계열을 이동시킨다.
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	// numerator 변수는 누적된 분자의 값을 저장하는데 사용되며
	// denominator 변수는 누적된 분모의 값을 저장하는 데 사용된다.
	var numerator float64
	var denominator float64

	// 자기상관(autocorrelation)의 각 항에 사용될
	// x값의 평균을 계산한다.
	xBar := stat.Mean(x, nil)

	// numerator(분자)를 계산한다.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// denominator(분모)를 계산한다
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}

package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	//"github.com/go-gota/gota/dataframe"
	//"github.com/montanaflynn/stats"
	//"github.com/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {

	// 관찰된 빈도를 정의한다.
	observed := []float64{
		260.0, // 정기적으로 운동하지 않는다고 답한 수를 나타낸다.
		135.0, // 가끔 운동한다고 답한 수를 나타낸다.
		105.0, // 정기적으로 운동한다고 잡한 수를 나타낸다.
	}

	// 관찰한 전체 수를 정의한다.
	totalObserved := 500.0

	// 관찰한 전체 수를 정의한다.
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// 카이젭곱 검정 통계량을 계산한다.
	chiSquare := stat.ChiSquare(observed, expected)

	// 표준 출력으로 검정 통계량을 출력한다.
	fmt.Printf("\nChi-square(카이제곱): %0.2f\n", chiSquare)

	// K의 자유도를 적용해 카이제곱 분포를 만든다.
	// 카이제곱 분포에 대한 자유도는 가능한 범주에서 1을 뺀 값이기 때문에
	// 이 경우 K=3-1=2 을 가진다.
	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	// 특정 검정 통계량에 대한 p-값을 계산한다.
	pValue := chiDist.Prob(chiSquare)

	// 표준 출력으로 P-값을 출력한다.
	fmt.Printf("p-value: %0.4f\n\n", pValue)
}

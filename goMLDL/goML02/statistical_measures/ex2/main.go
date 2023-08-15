package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/floats"
	"github.com/gonum/stat"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	//CSV 파일열기
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일에서 데이터프레임 생성하기
	irisDF := dataframe.ReadCSV(irisFile)

	// 이 변수에 대한 측정값을 확인하기 위해
	// "sepal_length" 열에서 float 값을 가져온다.
	petalLength := irisDF.Col("petal_length").Float() 
	 //sepal_lenthd은 값이 달라진다

	// 변수 최소값 계산하기
	minVal := floats.Min(petalLength)

	// 변수 최대값 계산하기
	maxVal := floats.Max(petalLength)

	// 변수의 중앙값을 계산한다.
	rangeVal := maxVal - minVal

	// 변수의 분산을 계산한다
	varianceVal := stat.Variance(petalLength, nil)

	// 변수의 표준 편차를 계산한다.
	stdDevVal := stat.StdDev(petalLength, nil)

	// 값을 정렬한다.
	inds := make([]int, len(petalLength))
	floats.Argsort(petalLength, inds)

	// 분위수를 계산한다.
	quant25 := stat.Quantile(0.25, stat.Empirical, petalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, petalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, petalLength, nil)

	// 표준 출력을 통해 결과를 출력한다.
	fmt.Printf("\n꽂받침 길이 요약 통계:\n")
	fmt.Printf("최대값: %0.2f\n", maxVal)
	fmt.Printf("최소값: %0.2f\n", minVal)
	fmt.Printf("범위값: %0.2f\n", rangeVal)
	fmt.Printf("분산: %0.2f\n", varianceVal)
	fmt.Printf("표준편차: %0.2f\n", stdDevVal)
	fmt.Printf("25 분위: %0.2f\n", quant25)
	fmt.Printf("50 분위: %0.2f\n", quant50)
	fmt.Printf("75 분위: %0.2f\n\n", quant75)
}
// 표준편차는 1.76이고 값의 전체 범위는 5.90 이다.분산과는 반대로 표준편차는
// 자체 값들과 단위가 동일하기 때문에 값의 범위에 따라 값이 달라지는 것을
// 볼 수 있다(표준편차 값은 전체 값 범위의 약 30%)

// 분위수를 살펴보자
// 25% 분위는 75%분위와 최대값 사이의 거리보다 최소값에 더 가깝다.따라서
// 분포에서 높은 값이 낮은 값보다 확산될 가능성이 크다는 것을 추론할 수 있다.

// 분포가 어떤 모양을 갖는지 정량화 하는데 도움을 얻기 위해 
// 중심 경향 측정 방법과 이런 측정값들을 조합해 사용하는 것도 가능하며
// 다른 방법도 존재한다. 
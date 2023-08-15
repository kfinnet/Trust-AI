package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/stat"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
)

func main() {

	// CSV 파일 열기
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일에서 데이터프레임 생성하기
	irisDF := dataframe.ReadCSV(irisFile)

	// "petal_length" 열에서 float값을 가져온다
	// 이 변수에 대한 측정값을 확인하기 위해
	petalLength := irisDF.Col("petal_length").Float()

	// 변수의 평균(Mean) 계산하기
	meanVal := stat.Mean(petalLength, nil)

	// 변수의 최빈값(Mode) 계산하기
	modeVal, modeCount := stat.Mode(petalLength, nil)

	// 변수의 중앙값(Median) 계산하기
	medianVal, err := stats.Median(petalLength)
	if err != nil {
		log.Fatal(err)
	}
	// 평균값, 최빈값, 중앙값이 모두 약간 다른 것을 볼 수 있다. 하지만
	// 평균값과 중앙값은 sepal_length 열에 있는 값과 매우 비슷한 것을 볼수 있다/
	// 표준 출력으로 결과 출력하기
	fmt.Printf("\n꽃받침 petalLenth 길이 요약 통계:\n")
	fmt.Printf("평균값: %0.2f\n", meanVal)
	fmt.Printf("최빈값: %0.2f\n", modeVal)
	fmt.Printf("최빈값 개수: %d\n", int(modeCount))
	fmt.Printf("중앙값: %0.2f\n\n", medianVal)
}
// 꽃잎 길이 petal_length 경우 평균값과 중앙값이 서로 비슷하지 않다.따라서
// 이미 이 정보로부터 데이터에 대한 통찰을 가질 수 있다. 평균값가 중앙값이 
// 비슷하지(가깝지) 않은 경우 높은 값과 낮은 값은 각각 평균값을 높게 또는
//낮게 끌어당긴다. 이런 영향은 중앙값에 서는 눈에 뛰지 않는다.
// 이런 현상을 기울어진 분포 skewed distribution 이라고 한다.
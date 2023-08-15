package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
)

func main() {

	// 붓꽃 데이터 집합을 읽고 golearn "인스턴스(instances)"로 설정한다.
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// 이 코드는 의사결정 트리를 제작하는데 포함되는
	// 랜덤 프로세스를 시작시킨다.
	rand.Seed(44111342)

	// 트리 당 2개의 수치를 갖는 10개의 트리로 구성된 랜덤 포레스트를 제작한다.
	// 이는 기본 설정으로 적당하다(일반적으로 트리당 수치는 sqrt(수치의 수)로 설정된다)
	// 앙상블
	rf := ensemble.NewRandomForest(10, 2) // 트리당 수치를 4로 올리면 정확도 증가

	// 5겹의 데이터 집합을 기반으로 모델을 성공적으로 훈련(학습)시키고 평가하기 위해
	// k-겹(k-fold) 교차 검증을 사용한다.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, rf, 5)
	if err != nil {
		log.Fatal(err)
	}

	// 교차 검증에 대한 정확도(accuracy)의 평균, 분산, 표준 편차를 구한다
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// 표준 출력을 통해서 교차 검증의 결과를 출력한다.
	fmt.Printf("\n정확도\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}

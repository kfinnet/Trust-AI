package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
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

	// 의사결정 트리 모델을 제작하기 위해 ID3 알고리즘을 사용한다. 또한
	// 훈련-가지의 분할을 조절하기 위한 매개변수 값을 0.6 으로 설정한다.
	tree := trees.NewID3DecisionTree(0.6)

	// 5겹의 데이터 집합을 기반으로 모델을 성공적으로 훈련(학습)시키고 평가하기 위해
	// k-겹(k-fold) 교차 검증을 사용한다.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, tree, 5)
	if err != nil {
		log.Fatal(err)
	}

	// 교차 검증에 대한 정확도(accuracy)의 평균, 분산, 표준 편차를 구한다.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// 표준 출력을 통해서 교차 검증의 결과를 출력한다.
	fmt.Printf("\n정확도\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}
// 91% 이상의 정확도를 얻는다. kNN모델보다 약간 나쁘지만 여전히 훌륭한 결과
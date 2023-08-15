package main

import (
	"fmt"
	"log"
	"math"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)
// github.com/sjwhitworth/golearn 제공모델 사용하려면, 데어터를 먼저
// 인스턴스라 불리는 golaern 패키지의 내부 포맷으로 변환해야한다.

func main() {

	// 붓꽃 데이터 집합을 읽고 golearn "인스턴스(instances)"로 설정한다
	irisData, err := base.ParseCSVToInstances("iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}
//다음은 kNN 모데을 초기화, 교차 검증을 빠르고 간단하게 수행해야한다.

	// 새 kNN 분류기를 초기화 한다. 간단한 유클리드 거리 측정법과
	// k=2 를 사용한다.
	knn := knn.NewKnnClassifier("euclidean", "linear", 2)

	// 5겹의 데이터 집합을 기반으로 모델을 성공적으로 훈련(학습)시키고 평가하기 위해
	// k-겹(k-fold) 교차 검증을 사용한다.
	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisData, knn, 5)
	if err != nil {
		log.Fatal(err)
	}

	// 교차 검증의 정확도(accuracy)에 대한 평균, 분산, 표준 편차를 구한다.
	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)
	log.Println(knn)
	//log.Println(cv)
    log.Println(mean)
	log.Println(stdev)
	// 표준출력을 통해서 교차 측정법의 결과를 출력한다.
	fmt.Printf("\n정확도\n%.2f (+/- %.2f)\n\n", mean, stdev*2)
}

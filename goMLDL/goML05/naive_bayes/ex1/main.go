package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func main() {

	// 훈련(학습) 데이터 집합을 읽고 golearn "인스턴스(instances)"로 설정한다
	trainingData, err := base.ParseCSVToInstances("training.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// 새로운 나이브 베이즈 분류기를 초기화 한다.
	nb := naive.NewBernoulliNBClassifier()

	// 나이브 베이즈 분류기를 적합(훈련)시킨다
	nb.Fit(convertToBinary(trainingData))

	// 대출 테스트 데이터 집합에서 데이터를 읽고
	// golearn "인스턴스(instances)"로 설정한다. 이번에는 테스트 집합의 형식을
	// 검증하기 위해 인스턴스의 이전 설정을 템플릿으로 활용한다.
	testData, err := base.ParseCSVToInstances("test.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// 예측을 수행한다.
	predictions, err := nb.Predict(convertToBinary(testData))

	// 혼동 행렬을 생성한다.
	cm, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}

	// 정확도를 계산한다.
	accuracy := evaluation.GetAccuracy(cm)
	fmt.Printf("\n정확도: %0.2f\n\n", accuracy)
}

// convertToBinary 함수는 golearn의 내장 기능을 사용해
// 데이터 집합의 레이블을 이진 헤이블 형식으로 변환한다.
func convertToBinary(src base.FixedDataGrid) base.FixedDataGrid {
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs {
		b.AddAttribute(a)
	}
	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}

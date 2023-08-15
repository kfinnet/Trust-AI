package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// 훈련(학습) 데이터 집합 파일을 연다.
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새 CSV reader를 생성한다.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// CSV records 를 모두 읽는다.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData와 labels 변수는 최종적으로 모델을 훈련시키는 데 사용될
	// float값을 저장하는 데 사용된다.
	featureData := make([]float64, 2*(len(rawCSVData)-1))
	labels := make([]float64, len(rawCSVData)-1)

	// featureIndex 변수는 수치를 저장하는 행렬 값의 현재 인덱스를
	// 추적하는 데 사용된다.
	var featureIndex int

	// 열을 순차적으로 이동하면서 슬라이스(slices)에 float값을 저장한다.
	for idx, record := range rawCSVData {

		// 헤더 열은 건너뛴다.
		if idx == 0 {
			continue
		}

		// FICO 점수 수치를 추가한다.
		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = featureVal

		// y절편을 추가한다.
		featureData[featureIndex+1] = 1.0

		// 수치 열에 대한 인덱스를 증가시킨다.
		featureIndex += 2

		// 클래스 레이블을 추가한다.
		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = labelVal
	}

	// 앞에서 저장한 수치로 행렬을 만든다.
	features := mat64.NewDense(len(rawCSVData)-1, 2, featureData)

	// 로지스틱 회귀분석 모델을 훈련(학습)시킨다.
	weights := logisticRegression(features, labels, 1000, 0.3)

	// 표준 출력을 통해 로지스틱 회귀분석 모델 공식을 출력한다.
	formula := "p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )"
	fmt.Printf("\n%s\n\nm1 = %0.2f\nm2 = %0.2f\n\n", formula, weights[0], weights[1])
}
// 모델을 훈련(학습)시키는 이 프로그램을 컴파일하고 실행하면
// 위와 같은 훈련(학습)된 로지스틱 회귀분석 공식 결과가 출력된다.

// logistic implements the logistic function, which
// is used in logistic regression.
func logistic(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// logisticRegression 함수는 주어진 데이터에 대해 
// 로지스틱 회귀분석 모델을 적합 fit (훈련)시킨다
func logisticRegression(features *mat64.Dense, labels []float64, numSteps int, learningRate float64) []float64 {

	// 가중치를 임의로 추기화 한다.
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for idx, _ := range weights {
		weights[idx] = r.Float64()
	}

	// 가중치를 반복해서 최적화 한다.
	for i := 0; i < numSteps; i++ {

		// 이번 반복 실행에 대한 오차를 누적시키기 위한 변수를 초기화 한다.
		var sumError float64

		// 각 레이블에 대한 예측을 수행하고 오차를 누적시킨다.
		for idx, label := range labels {

			// 이 레이블에 해당하는 값을 가져온다.
			featureRow := mat64.Row(nil, idx, features)

			// 이번 반복 실행의 가중치에 대한 오차를 계산한다.
			pred := logistic(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			// 가중치값을 업데이트 한다.
			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}
	return weights
}
// 로지스틱 회귀분석 모델을 훈련(학습) 데이터 집합을 통해 훈련시키기 위해
// encoding/csv 사용해 훈련(학습) 데이터 파일을 읽어온 다음,
// logisticRegression 함수에 필요한 매개변수를 전달한다.
// 이 과정은 다음과 같으며 훈련(학습)된 로지스틱 회귀분석 공식을 표준출력을 통해 출력
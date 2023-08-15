package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

func main() {

	// 훈련(학습) 데이터 집합 파일을 연다
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새 CSV reader를 생성한다.
	reader := csv.NewReader(f)

	// CSV 레코드를 모두 읽는다.
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 여기에서는 TV 수치와 y절편에 대해 판매량(y)을 모델링하려고 한다.
	// 따라서 github.com/sajari/regression을 사용해 모델을 훈련(학습)시키기 위해 필요한
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// 루프를 통해 CSV에서 레코드를 읽고 regression 값에 훈련(학습) 데이터를 추가한다.
	for i, record := range trainingData {

		// 헤더는 건너뛴다
		if i == 0 {
			continue
		}

		// 판매량 값을 구문 분석해 읽는다.
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// TV 값을 구문 분석해 읽는다.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 이 값들을 regression 값에 추가한다.
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	// 회귀분석 모델을 훈련(학습)/적합한다.
	r.Run()

	// 훈련(학습)을 거친 모델 매개변수를 출력한다.
	fmt.Printf("\nRegression Formula 회귀분석계산:\n%v\n\n", r.Formula)

	// 데스크 데이터 집합 파일을 연다.
	f, err = os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린파일을 읽는 새 CSV reader 를 생성한다.
	reader = csv.NewReader(f)

	// 모든 CSV records를 읽는다
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 루프를 통해 y를 예측하는 테스트 데이터를 읽고
	// 평균 절대값 오차를 활용해 예측된 수치를 평가한다.
	var mAE float64
	for i, record := range testData {

		// 헤더는 건너뛴다.
		if i == 0 {
			continue
		}

		// 관찰된 판매량 또는 "y" 값을 구문 분석해 읽는다.
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// TV값을 구문 분석해 읽는다.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 훈련된 모델을 사용해 예측을 수행한다.
		yPredicted, err := r.Predict([]float64{tvVal})

		// 평균 절대 오차(MAE)에 추가한다.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// 표준 출력으로 MAE를 출력한다.
	fmt.Printf("MAE평균절대오차값 = %0.2f\n\n", mAE)
}

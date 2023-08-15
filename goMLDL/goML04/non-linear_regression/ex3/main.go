package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	// 테스트 데이터 집합 파일을 연다
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새 CSV reader를 생성한다.
	reader := csv.NewReader(f)

	// 모든 CSV 레코드를 읽는다.
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

		// 관촬된 판매량 또는 "y" 값을 구문 분석해 읽는다.
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// TV값을 구문 분석해 읽는다.
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Radio 값을 구문 분석해 읽는다
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Newspaper 값을 구문 분석해 읽는다
		newspaperVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 훈련된 모델을 사용해 예측을 수행한다.
		yPredicted := predict(tvVal, radioVal, newspaperVal)

		// 평균 절대 오차(MAE)에 추가한다.
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// 표준 출력으로 MAE를 출력한다.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)
}

// predict 함수는 훈련된 회귀분석 모델을 사용해 예측을 수행한다.
// TV, Radio, and Newspaper value.
func predict(tv, radio, newspaper float64) float64 {
	return 3.038 + tv*0.047 + 0.177*radio + 0.001*newspaper
}

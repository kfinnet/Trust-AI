package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)
// predict 함수를 사용하면 앞부분 소개한 평가 방법중 하나를 사용해
// 훈련된 회귀분석 모델을 평가할 수 있다. 여기서는 다음 코드와 같이
// 정확도 accuracy를 사용해 보자
func main() {

	// 데스트 데이터 집합 파일을 연다
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새 CSV reader를 생성한다.
	reader := csv.NewReader(f)

	// Observed 와 predicted 변수는 레이블이 지정된 데이터 파일로부터 읽어온
	// 관찰값 및 예측값을 저장하는 데 사용된다.
	var observed []float64
	var predicted []float64

	// line 변수는 로그를 위해 열의 수를 추척하는 데 사용된다.
	line := 1

	// 열에서 예기치 않은 유형을 찾기 위해 레코드를 읽는다.
	for {

		// 열을 읽는다. 파일에 끝 부분에 도달했는지 확인한다.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 헤더는 건너뛴다.
		if line == 1 {
			line++
			continue
		}

		// 관찰값을 읽는다.
		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// 필요한 예측을 수행한다.
		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal := predict(score)

		// 기대하는 유형인 경우, 해당 레코드를 슬라이스(slice)에 추가한다.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// 이 변수는 true positive와 true negatice값의 횟수를 
	// 저장하는데 사용된다.
	var truePosNeg int

	// true positive/negative 횟수를 누적시킨다.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// 정확도 (accuracy)를 계산한다(부분집합 정확도)
	accuracy := float64(truePosNeg) / float64(len(observed))

	// 표준 출력을 통해 저오학도(Accuracy)값을 출력한다
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}

// predict 함수는 훈련(학습)된 로지스틱 회귀분석 모델을 기반으로
// 예측을 수행한다.
func predict(score float64) float64 {

	// 예측 확률을 계산한다.
	p := 1 / (1 + math.Exp(-13.65*score+4.89))

	// 해당하는 클래스를 출력한다.
	if p >= 0.5 {
		return 1.0
	}

	return 0.0
}

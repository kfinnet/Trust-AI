package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
     // 이진 관찰값과 예측값을 연다
	f, err := os.Open("labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 앞에서 연 파일을 읽는 새 CSV reader를 생성한다.
	reader := csv.NewReader(f)

	// Oberved 및 predicted 변수는 labeled 데이터 파일에서 읽어온
	// 관찰값과 예측값을 저장하는 데 사용된다.
	var observed []int
	var predicted []int

	// line 변수는 로그를 위해 행(row)의 수를 기록한다.
	line := 1

	// 열(colums)에서 예기치 않은 유형에 대한 레코드를 읽는다.
	for {

		// 행을 읽어온다. 파일 끝에 도달했는지 확인한다.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 헤더는 건너뛴다.(생략)
		if line == 1 {
			line++
			continue
		}

		// 관찰된 값과 예측값을 읽는다.
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// 예상되는 유형인 경우 슬라이스에 해당 레코드를 추가한다.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// classes 변수는 labeled 데이터에서 3개의 요소값 클래스를 포함한다.
	classes := []int{0, 1, 2}

	// 각 클래쓰에 대해 루프를 통해서 순회 작업한다.
	for _, class := range classes {

		// 이 변수들은 true positives 횟수와
		// false positives/false negative 횟수를 저장하는 데 사용된다.
		var truePos int
		var falsePos int
		var falseNeg int

		// true positive와 false positive의 횟수를 누적시킨다.
		for idx, oVal := range observed {

			switch oVal {

			// 관찰된 값이 특정 클래스인 경우 
			// 예측한 값이 해당 클래스였는지를 확인한다.
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}

				falseNeg++

			// 관찰된 값이 다른 클래스인 경우
			// 예측값이 false positive 였는지 확인한다.
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		// 정밀도(precision)를 계산한다.
		precision := float64(truePos) / float64(truePos+falsePos)

		// 재현율(recall)을 계산한다.
		recall := float64(truePos) / float64(truePos+falseNeg)

		// 표준출력을 통해서 정밀도와 재현율을 출력한다.
		fmt.Printf("\nPrecision정밀도 (class클래스 %d) = %0.2f", class, precision)
		fmt.Printf("\nRecall재현율 (class클래스 %d) = %0.2f\n\n", class, recall)
	}
}

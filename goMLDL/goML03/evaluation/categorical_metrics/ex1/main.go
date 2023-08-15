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
		if err == io.EOF { //End of file
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

	// 이 변수는 true positive 와 negative 값의 횟수를
	// 저장하는데 사용된다.
	var truePosNeg int

	// true positive/negative 횟수를 누적시킨다.
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++ //자동 증가 누적시킨다
		}
	}

	// 정확도(accuracy)를 계산한다(부분 정확도) 
	// accuracy : (TP + TN) / (TP + TN + FP + FN)
	accuracy := float64(truePosNeg) / float64(len(observed)) //len( )

	// 표준 출력을 통해 정확도값을 출력한다.
	fmt.Printf("\nAccuracy 정확도 = %0.2f\n\n", accuracy)
}

package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 렌딩클럽 대출 데이터 집합 파일을 연다.
	f, err := os.Open("loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새로운 CSV reader 를 생성한다.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	// CSV 레코드를 모두 읽는다.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 출력 파일을 생성한다.
	f, err = os.Create("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV writer를 생성한다.
	w := csv.NewWriter(f)

	// 열을 순차적으로 이동하면서 구문 분석된 값을 쓴다.
	for idx, record := range rawCSVData {

		// 헤더 열은 건너 뛴다.
		if idx == 0 {

			// 헤더를 출력 파일에 쓴다.
			if err := w.Write([]string{"FICO_score", "class"}); err != nil {
				log.Fatal(err)
			}
			continue
		}

		// 읽어온 값을 저장하기 위한 슬라이스(slice)를 초기화한다.
		outRecord := make([]string, 2)

		// FICO 점수를 구문 분석하고 표준화한다.
		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score-640.0)/(830.0-640.0), 'f', 4, 64)

		// 이자율 클래스를 구문 분석한다.
		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"

			// 레코드를 출력 파일에 쓴다.
			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord[1] = "0.0"

		// 레코드를 출력 파일에 쓴다.
		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// 버퍼에 저장된 데이터를 기본 writer(표준 출력)에 쓴다.
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

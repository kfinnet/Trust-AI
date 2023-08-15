package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord는 CSV 파일에서 성공적으로 읽어온 행을 포함저장한다.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	//레코드를 루프로 처리하기 전에 이 값들을 초기화 한다

	// CSV에서 성공적으로 읽어온 레코드를 저장하는 값을 생성한다.
	var csvData []CSVRecord

	// line will help us keep track of line number for logging.
	line := 1

	// 레코드를 읽고 예상하지 못한 타입을 찾는다.
	for {

		// 열을 읽는다. 파일 종료 지점에 도달했는지 확인한다.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 열을 저장하기 위한 CSVRecord를 생성한다.
		var csvRecord CSVRecord

		// 기대하는 타입을 기반으로 레코드의 각 값을 읽는다.
		for idx, value := range record {

			// 문자열 행에 대해 문자열로 레코드의 값을 읽는다.
			if idx == 4 {

				//  값이 빈 문자열이 아닌지 확인한다.
				// 해당 값이 빈 문자열인 경우
				// 구문 분석을 처리하는 루프를 중단한다.
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// CSVRecord에 문자열값을 추가한다.
				csvRecord.Species = value
				continue
			}

			// 문자열 행이 아닌 경우 레코드의 값을 float64로 읽는다.
			var floatValue float64

			// 레코드의 값이 float로 읽혀지지 않으면 로드에 기록하고
			// 구문 분석 처리 루프를 중단한다.
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// CSVRecord의 해당 필드에 float값을 추가한다.
			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		// 앞에서 생성해둔 csvData에 성공적으로 읽어온 레코드를 추가한다
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		// Increment the line counter.
		line++
	}

	fmt.Printf("successfully parsed %d lines\n", len(csvData))
}

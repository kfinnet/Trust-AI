package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// 훈련(학습) 데이터 집합 파일을 연다.
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 열린 파일을 읽는 새 CSV reader 를 생성한다.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	// CSV 레코드를 모두 읽는다.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData는 최종적으로 수치를 나타내는 행렬을 만드는데 사용될
	// 모든 float값을 저장한다.
	featureData := make([]float64, 4*len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	// featureData 및 float값의 현재 인덱스를 추적하는 데 사용된다.
	var featureIndex int
	var yIndex int

	// 순차적으로 행을 읽어 float값을 저장하는 슬라이스(slice)에 추가한다
	for idx, record := range rawCSVData {

		// 헤데 행을 건너뛴다.
		if idx == 0 {
			continue
		}

		// 루프를 통해 float 열을 읽는다.
		for i, val := range record {

			// 값을 float로 변환한다.
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("Could not parse float value")
			}

			if i < 3 {

				// 모델에 y절편을 추가한다.
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				// float값을 저장하는 슬라이스(slice)에 float값을 추가한다
				featureData[featureIndex] = valParsed
				featureIndex++
			}

			if i == 3 {

				// y float값을 저장하는 슬라이스(slice)에 float값을 추가한다.
				yData[yIndex] = valParsed
				yIndex++
			}

		}
	}

	// 희귀분석 모델에 입력될 행렬들을 만든다.
	features := mat64.NewDense(len(rawCSVData), 4, featureData)
	y := mat64.NewVector(len(rawCSVData), yData)

	if features != nil && y != nil {
		fmt.Println("Matrices formed for ridge regression")
	}
}

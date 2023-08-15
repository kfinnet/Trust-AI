package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type centroid []float64

func main() {

	// CSV 파일을 연다
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV파일로 부터 dataframe 을 생성한다.
	irisDF := dataframe.ReadCSV(irisFile)

	// CSV 파일에 포함되어 있는 세 개의 개별 품종의 이름을 정의한다.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// 중심 정보를 저장할 맵(map)을 생성한다.
	centroids := make(map[string]centroid)

	// 데이터 집합을 세 개의 개별 dataframe 으로 필터링 한다.
	// 각 dataframe 붓꽃 품종 중의 하나에 해당한다.
	for _, species := range speciesNames {

		// 원본 데이터 집합을 필터링한다.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// 수치들의 평균을 계산한다.
		summaryDF := filtered.Describe()

		// 각 차원의  평균을 해당하는 중심에 입력한다.
		var c centroid
		for _, feature := range summaryDF.Names() {

			// 관련없는 열은 건너뛴다
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// 이 중심 정보를 맵(map)에 추가한다.
		centroids[species] = c
	}

	// 데이터를 확인하기 위해 중심을 출력한다.
	for _, species := range speciesNames {
		fmt.Printf("%s centroid: %v\n", species, centroids[species])
	}
}

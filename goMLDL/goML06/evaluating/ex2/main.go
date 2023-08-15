package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/floats"
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

	// CSV 파일로부터 dataframe을 생성한다.
	irisDF := dataframe.ReadCSV(irisFile)

	// CSV 파일에 포함되어 있는 세 개의 개별 품종의 이름을 정의한다.
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// 중심 정보를 저장할 맵(map)을 생성한다.
	centroids := make(map[string]centroid)

	// 각 군집에 대해 필터링된 dataframe 저장하는 맵(map)을 생성한다.
	clusters := make(map[string]dataframe.DataFrame)

	// 데이터 집합을 세 개의 개별 dataframe으로 필터링 한다.
	// 각 dataframe 붓꽃 품종 중의 하나에 해당한다.
	for _, species := range speciesNames {

		// 원본 데이터 집합을 필터링한다.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)

		// 필터링을 거친 dataframe을 군집을 저장하는 map에 추가한다.
		clusters[species] = filtered

		// 수치들의 평균을 계산한다.
		summaryDF := filtered.Describe()

		// 각 차원의 평균을 해당하는 중심에 입력한다.
		var c centroid
		for _, feature := range summaryDF.Names() {

			// 관련 없는 열은 건너뛴다.
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// 이 중심 정보를 aoq(map)에 추가한다
		centroids[species] = c
	}

	// 레이블을 문자열 배열로 변환하고 편의를 위해
	// float 열의 이름을 저장하는 배열을 생성한다.
	labels := irisDF.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	// 루프를 통해 레코드를 읽고 실루엣 계수의 평균을 누적시킨다.
	var silhouette float64

	for idx, label := range labels {

		// a 변수는 a에 대한 누적 값을 저장하는데 사용된다.
		var a float64

		// 루프를 통해 동일한 군집 내의 데이터 요소를 읽는다.
		for i := 0; i < clusters[label].Nrow(); i++ {

			// 비교를 위한 데이터 요소를 얻는다.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			// a에 추가한다.
			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())
		}

		// 가장 가까운 다른 군집을 구한다.
		var otherCluster string
		var distanceToCluster float64
		for _, species := range speciesNames {

			// 동일한 데이터를 갖는 군집은 건너뛴다.
			if species == label {
				continue
			}

			// 현재 클러스터에서 해당 군집 사이의 거리를 계산한다.
			distanceForThisCluster := floats.Distance(centroids[label], centroids[species], 2)

			// 필요한 경우 현재 군집을 교체한다.
			if distanceToCluster == 0.0 || distanceForThisCluster < distanceToCluster {
				otherCluster = species
				distanceToCluster = distanceForThisCluster
			}
		}

		// b 변수는 b에 대한 누적 값을 저장하는 데 사용된다.
		var b float64

		// 루프를 통해 가장 가까운 다른 군집 내의 데이터 요소를 읽는다.
		for i := 0; i < clusters[otherCluster].Nrow(); i++ {

			// 비교를 위해 데이터 요소를 얻어온다.
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)

			// b에 추가한다.
			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())
		}

		// 평균 실루엣 계수레 추가한다.
		if a > b {
			silhouette += ((b - a) / a) / float64(len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}

	// 표준 출력을 통해 최종 평균 실루엣 계수를 출력한다.
	fmt.Printf("\n평균 실루엣 계수: %0.2f\n\n", silhouette)
}
// dataframe.DataFrame 으로 floats값을 얻기 위한 편의 함수를 생성한다.
// dfFloatRow 함수는 주어진 인덱스와 주어진 열의 이름을 사용해
// DataFrame으로 부터 float값의 배열을 얻어온다.
func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row
}

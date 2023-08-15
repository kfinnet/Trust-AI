package main

import (
	"fmt"
	"log"
	"os"
	"gonum.org/v1/gonum/floats"
	//"github.com/gonum/floats"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// 운전자 데이터 집합 파일을 연다
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV 파일로부터 dataframe 을 생성한다.
	driverDF := dataframe.ReadCSV(f)

	// Distance_Feature에 대한 열의 값을 추출한다.
	distances := driverDF.Col("Distance_Feature").Float()

	// clusterOne and clusterTwo 변수는 도표에 대한 값을 저장하는 데 사용된다.
	var clusterOne [][]float64
	var clusterTwo [][]float64

	// clusterOne과 clusterTwo 변수를 데이터로 채운다.
	for i, speed := range driverDF.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{distances[i], speed}, []float64{50.05, 8.83}, 2)
		distanceTwo := floats.Distance([]float64{distances[i], speed}, []float64{180.02, 18.29}, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{distances[i], speed})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{distances[i], speed})
	}

	// 군집 평가 측정치를 출력한다.
	fmt.Printf("\nCluster군집 1 Metric측정수치: %0.2f\n", withinClusterMean(clusterOne, []float64{50.05, 8.83}))
	fmt.Printf("\nCluster군집 2 Metric측정수치: %0.2f\n", withinClusterMean(clusterTwo, []float64{180.02, 18.29}))
}

// withinClusterMean 함수는 군집의 여러 위치들과 해당 군집의 중심 위치 사이의
// 평균 거리를 계산한다.
func withinClusterMean(cluster [][]float64, centroid []float64) float64 {

	// meanDistance 변수는 결과를 저장하는 데 사용된다.
	var meanDistance float64

	// 루프를 통해 군집 내 요소를 읽고 평균 거리를 계산한다.
	for _, point := range cluster {
		meanDistance += floats.Distance(point, centroid, 2) / float64(len(cluster))
	}

	return meanDistance
}

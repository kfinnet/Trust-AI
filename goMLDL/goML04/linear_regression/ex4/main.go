package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	// 광고 데이터 집합 파일을 연다.
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV 파일로부터 dataframe 생성한다.
	// 열의 유형은 추론된다.
	advertDF := dataframe.ReadCSV(f)

	// 각 집합에서 요소의 수를 계산한다.
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum+testNum < advertDF.Nrow() {
		trainingNum++
	}

	// 훈련(학습)용 인덱스와 테스트용 인덱스를 저장할 배열을 생성한다.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 훈련(학습)용 인덱스를 배열에 저장한다.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// 테스트용 인덱스를 배열에 저장한다.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// 각 데이터 집합에 대한 데이터프레임을 생성한다.
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)

	// 데이터를 파일에 쓸 때 사용될 맵(map)을 생성한다.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// 각각의 파일을 생성한다.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// 필터링을 거친 데이터 집합 파일을 저장한다.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// buffered writer를 생성한다.
		w := bufio.NewWriter(f)

		// 데이터프레임을 CSV 파일로 쏜다.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
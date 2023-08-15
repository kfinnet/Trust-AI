package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

// 프로그램을 컴파일하고 실행하면 훈련(학습) 및 테스트 예제를 포함하는
// 두개의 파일이 생성된다.
func main() {

	// 정리된 대출 데이터 집합 파일을 연다
	f, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV 파일로 부터 dataframe을 생성한다.
	// 열의 유형을 추론된다.
	loanDF := dataframe.ReadCSV(f)

	// 각 집합에서 항목의 수를 계산한다.
	trainingNum := (4 * loanDF.Nrow()) / 5
	testNum := loanDF.Nrow() / 5
	if trainingNum+testNum < loanDF.Nrow() {
		trainingNum++
	}

	// 훈련(학습) 데이터 집합과 테스트 데이터 집합에서 사용할 인덱스를 생성한다.
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 루프를 통해 훈련(학습)용 인덱스를 저장한다.
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// 루프를 통해 테스트용 인덱스를 저장한다.
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// 훈련(학습) 및 테스트용 dataframe을 생성한다.
	trainingDF := loanDF.Subset(trainingIdx)
	testDF := loanDF.Subset(testIdx)

	// 데이터를 파일에 쓸 때 사용할 맵(map)을 생성한다.
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// 파일을 각 각 생성한다.
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// 필터링을 거친 데이터 집합 파일을 저장한다.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// 버퍼 Writer를 생성한다.
		w := bufio.NewWriter(f)

		// dataframe 을 CSV 파일로 쓴다.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}

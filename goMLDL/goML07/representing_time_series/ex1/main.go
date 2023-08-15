package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

// representation : 전체 데이터를 단순화해서 패턴을 살펴보는 방법

func main() {

	// CSV 파일을 연다
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV 파일로 부터 dataframe을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// 검사를 위해 표준출력을 통해 레코드를 표시한다.
	// Gota는 깔끔한 출력을 위해 dataframe의 형식을 지정한다.
	fmt.Println(passengersDF)
}

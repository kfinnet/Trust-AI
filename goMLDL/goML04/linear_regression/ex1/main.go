package main

import (
	"fmt"
	"log"
	"os"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	// CSV file을 연다
	advertFile, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	// CSV file로 부터 dataframe을 생성한다.
	advertDF := dataframe.ReadCSV(advertFile)

	// Describe 메소드를 사용해
	// 모든 열에 대한 요약 통계를 한번에 계산한다.
	advertSummary := advertDF.Describe()
    // Describe prints the summary statistics for each column of the dataframe
	// 표준 출력을 통해 요약 통계를 출력한다
	fmt.Println(advertSummary)
}

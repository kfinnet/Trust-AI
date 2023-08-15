package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// CSV열기
	f, err := os.Open("myfile.csv")
	if err != nil {
		log.Fatal(err)
	}

	// CSV 레코드에서 읽기
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 정수열에서 최댓값 구하기
	var intMax int
	for _, record := range records {

		// 정수값 해석하기
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// 적당한 경우 최댓값 바꾸기
		if intVal > intMax {
			intMax = intVal
		}
	}

	// 최대값 출력하기
	fmt.Println(intMax)
}

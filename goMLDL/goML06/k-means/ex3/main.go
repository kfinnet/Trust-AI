package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	//"github.com/mdesenfants/gokmeans"
	"github.com/mash/gokmeans"
)

func main() {

	// 운전자 데이터 집합 파일을 연다.
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader를 생성한다.
	r := csv.NewReader(f)
	r.FieldsPerRecord = 3

	// 입력 데이터를 저장하기 위한 gokmeans.Node의 배열을
	// 초기화 한다
	var data []gokmeans.Node

	// gokmeans.Node의 배열에 값을 저장하기 위해
	// 루프를 통해 레코드를 읽는다.
	for {

		// 레코드를 읽고 오류를 확인한다.
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// 헤러는 건너뛴다.
		if record[0] == "Driver_ID" {
			continue
		}

		// point 변수를 초기화한다.
		var point []float64

		// point 변수에 데이터를 저장한다.
		for i := 1; i < 3; i++ {

			// float 값을 구문 분석해 값을 읽는다.
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(err)
			}

			// point 변수에 이 값을 추가한다.
			point = append(point, val)
		}

		// 데이터에 point 변수를 추가한다.
		data = append(data, gokmeans.Node{point[0], point[1]})
	}
    // 군집의 생성은 gomeans.Train(...) 함수를 호출해 쉽게 생성할수 있다
	// 특히 k = 2 에서 최대 50번 반복시켜 이 함수를 호출한다.
	// k-평균 클러스터링을 사용해 군집을 생성한다.
	success, centroids := gokmeans.Train(data, 2, 50)
	if !success {
		log.Fatal("Could not generate clusters")
	}

	// 표준 출력으로 중심 위치들을 출력한다.
	fmt.Println("The centroids for our clusters are 이군집의 중심 위치들:")
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}

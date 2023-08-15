package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
  // 이 버전의 iris.csv 파일에는 필드가 추가된 행이있다.
  // 각 레코드는 5개의 필드를 가져야 한다는 점을 알고 있기 때문에
  // reader.FieldsPerRecord값을 5로 설정한다.

		// iris 데이터셋 파일을 연다
	f, err := os.Open("../data/iris_unexpected_fields.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// 각 라인마다 5개의 필드가 있어야 한다
	// FieldsPerRecord를 5로 설정하면 CSV의 각 행에 대한 정확한
	// 필드 수가 있는지 확인할 수 있다
	reader.FieldsPerRecord = 5

	// 이렇게 하면 CSV파일에서 레코드를 읽을수 있으며 
    // 예상하지 못한 필드를 확인할 수 있고 이를 통해서
	// 데이터(레코드 열, 행의)무결성을 유지할 수 있다.

	// rawCSVData는 성공적으로 파싱(구문해석된) 행을 저장한다.
	var rawCSVData [][]string

	// 레코드를 읽고 예상하지 못한 필드 수를 찾는다
	for {

		// 열을 읽는다. 파일 종료 지점에 도달했는지 확인한다.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 값을 읽는 과정에서 오류가 발생하면 
		// 오류를 로그에 기록하고 계속 진행한다
		if err != nil {
			log.Println(err)
			continue
		}

		// 레코드가 기대한 필드 수를 갖는 경우
		// 데이터 집합에 해당 레코드를 추가한다
		rawCSVData = append(rawCSVData, record)
	}
	fmt.Printf("parsed %d lines successfully\n", len(rawCSVData))
}
// 여기는 오류를 처리하는 방법으로 로그에 기록하는 방식을 선택
// 그리고 성공적으로 읽어온 레코드는 rawCSVData에 수집한다.
// 다양한 방법으로 이런 오류를 처리 할수 있다.
// 중요한 점은 데이터의 예상되는 특성을 확인해 
// 어플리케이션의 데이터의 무결성을 높였다는 점이다.
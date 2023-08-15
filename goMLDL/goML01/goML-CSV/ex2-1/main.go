package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// iris 데이터셋 파일을 연다
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
    // 열린 파일을 읽어 오는 새 CSV reader를 생성한다
	// 그러면 CSV파일의 모든 레코드(행)를 읽을수 있고
	// 이 레코드들은 [][]string 형태로 임포트(import)된다.
	reader := csv.NewReader(f)
   // 라인당 필드 수를 모른다고 가정한다.  
	// FieldsPerRecord를 음수로 설정해서
	// 각 행의 필드의 수를 얻을 수 있다.
	reader.FieldsPerRecord = -1

	// rawCSVData는 성공적으로 읽어온 행의 데이터를 저장한다.
	var rawCSVData [][]string

	// 레코드를 하나씩 읽는다
	for {

		// 열을 읽는다. EOF : End of File 파일 종료 지점에 도달했는지 확인한다.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 데이터 집합에 레코드를 추가한다
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}

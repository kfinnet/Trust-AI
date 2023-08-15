package main

import (
	"encoding/csv"
	"fmt"
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
	// FieldsPerRecord(필드수를 알아낸다)를 음수로 설정(필드 시작을 0부터 최대 양수값까지)해서
	// 각 행의 필드의 수를 얻을 수 있다.
	reader.FieldsPerRecord = -1

	// 모든 CSV레코드를 읽는다
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
    // 또한 무한 루프에서 레코드를 한 번에 하나씩 읽을 수도 있다.
	// 이를 위해서는 데이터를 모두 읽은 후에 루프를 종료할 수 있도록
	// 파일 종료 지점(io.EOF)을 확인해야 한다.

	fmt.Println(rawCSVData)
}

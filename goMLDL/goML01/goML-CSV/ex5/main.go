package main

import (
	"fmt"
	"log"
	"os"
	"github.com/go-gota/gota/dataframe"
)
//dataframe 패키지를 사용하면 표만들기,재조합 집합 정렬시킨 효과를 볼수 있다
func main() {

	// CSV 파일열기
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일로부터 데이터프레임 생성하기
	// 열의 유형은 추론된다
	irisDF := dataframe.ReadCSV(irisFile)
	// 반환된 포인터를 dataftame.ReadCSV() 함수에 제공

	// 검사를 위해 레코드를 stdout(표준 출력)으로 보여준다
	// Gota 패키지는 적절한 형태로 출력될 수 있도록 
	// 데이터프레임 형식을 지정
	fmt.Println(irisDF)
}

// 위 프로그램을 실행, 구문 분석 과정에서 추론된 형식이 지정된 데이터가 
// 적절한 형태로 출력되는 것을 볼수 있다

// 일단 데이터를 데이터 프레임으로 읽고 나면 필터링, 
// 부분 집합으로 나누기
// 데이터 선택을 쉽게 할 수 있다
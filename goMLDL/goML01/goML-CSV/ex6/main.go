package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

// 일단 데이터를 데이터프레임으로 읽고 나면 필터링,
// 부분집합 나누기, 데이터 선택을 쉽게 할수 있다

func main() {

	//CSV파일을 가져온다
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

  	// CSV 파일로부터 데이터프레임 생성하기
	// 열의 유형은 추론된다
	irisDF := dataframe.ReadCSV(irisFile)

	// 데이터 프레임의 필터를 생성한다
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// 붓꽃(iris) 품종이 "Iris-versicolor"인 행만 볼수 있도록
	// 데이터프레임을 필터링한다
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// Output the results to standard out.
	fmt.Println(versicolorDF)

	// 데이터프레임을 다시 필터링한다. 하지만 이번에는 
	// sepal_width and species 열만 선택한다
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_length", "species"})
	fmt.Println(versicolorDF)

	// 데이터프레임을 필터링하고 다시 선택한다, 하지만 이번에는
	// 처음 세계의 결과만 보여준다
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_length", "species"}).Subset([]int{2, 3, 4})
	fmt.Println(versicolorDF)

}

// 위 내용은 github.com/go-gota/gota/dataframe 패키지의
// 극히 일부만 살펴본 것.
// 데이터 집합 병합, 다른 포맷으로 출력,
// JSON 데이터를 처리하는 것도 가능
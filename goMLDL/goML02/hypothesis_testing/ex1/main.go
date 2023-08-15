package main

import (
	"fmt"

	"github.com/gonum/stat"
)

func main() {

	// 관찰값과 예상값을 정의한다.
	// 이 값들은 대부분 사용자 데이터로부터 얻어진다.
	// (웹 사이트 방문 등).
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// 카이제곱 검정 통계량을 계산한다.
	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Println(chiSquare)
}

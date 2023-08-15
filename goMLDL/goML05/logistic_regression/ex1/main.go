package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(logistic(1.0))
}

// logistic 함수는 로지스틱 함수를 구현하며
// 로지스틱 회귀분석에 사용된다
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

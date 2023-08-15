package main

import (
	"fmt"

	"github.com/gonum/floats"
)

func main() {

	// Distance 함수의 마지막 인자를 통해 지정한 유클리드 거리를 계산한다.
	distance := floats.Distance([]float64{1, 2}, []float64{3, 4}, 2)

	fmt.Printf("\nDistance: %0.2f\n\n", distance)
}

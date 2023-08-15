package main

import (
	"fmt"

	"github.com/gonum/floats"
)

func main() {

	// 슬라리스로 표현되는 두 "벡터"를 초기화 한다.
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// A와 B 벡터의 내적을 계산한다.
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// A 벡터의 각 요소에 1.5 를 곱한다.
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// B 벡터의 놈(norm)/ 길이를 계산한다.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}

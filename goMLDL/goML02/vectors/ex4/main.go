package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/matrix/mat"
)
// gonum.org/v1/gonum/mat 를 활용해 비슷한 연산을 작업할 수 있다.

func main() {

	// 슬라이스로 표현되는 두 "벡터"를 초기화 한다.
	vectorA := mat64.NewVector(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat64.NewVector(3, []float64{-7.2, 4.2, 5.1})

	// A와 B 벡터의 내적을 계산한다.
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := mat64.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// A 벡터의 각 요소에 1.5 를 곱한다.
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// B 벡터의 놈(norm)/ 길이를 계산한다.
	normB := blas64.Nrm2(3, vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
// 두 경우의 의미는 비슷. 
// 벡터(행렬이 아닌)로만 작업하는 경우에는 float로 이루어진 슬라이스에 대해
// 좀 더 빠르고 가벼운 연산이 필요한데, 
// gonum.org/v1.gonum/floats  좋은 선택이 될 수 있다.
// 하지만 행렬과 벡터 모두를 활용해 작업하는 경우에는 
// 벡터/행렬에 대한 좀 더 넓은 범위의 기능에 접근해야 하기 때문에 
// gonum.org/v1/gonum/mat 를 사용하는 것이 더 나은 선택이다.
// gonum.org/v1/gonum/blas/blas64 를 사용하는 것이 좋은 선택인 경우도 있다 
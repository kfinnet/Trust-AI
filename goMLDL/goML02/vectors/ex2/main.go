package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)
// gonum은 float64로 구성된 슬라이스를 계산하기 위한
// gonum.org/v1/gonum/floats  를 제공하고 벡터 유형(관련 메소드 포함)과
// 행렬을 활용한 계산이 가능한 
// gonum.org/v1/gonum/mat 를 제공한다.

func main() {

	// 새로운 벡터값 생성하기
	myvector := mat64.NewVector(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}
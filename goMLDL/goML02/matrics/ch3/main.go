package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// 새 행렬 a를 생성한다.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// 행렬의 전치를 계산하고 이를 출력한다.(*전치계산을 하면 역행렬처리로 변환)
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// 행렬 a의 행렬식을 계산하고 이를 출력한다.
	// det(a) =(1x4x6=)24+(0x0x3=)0 +(0x2x5=)0
	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// 행렬 a의 역행렬을 구하고 이를 출력한다.
	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat64.Formatted(aInverse, mat64.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}
// 무결성과 가독성을 유지위해, Go 명시적 오류처리 기능활용
// 모든행렬이 역행렬을 갖지는 않는다.
// 즉, 역행렬을 구할 수 없는 행렬도 존재한다.
// 행렬 규모가 큰 데이터 집합 작업시 다양한 상황발생 할수 있고
// 이런 상황에서도 우리가 제작하는 프로그램이 예상대로 동작할 수 있도록 
// 오류 처리가 중요하다.
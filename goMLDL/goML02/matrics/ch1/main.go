package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)
// gonum.org/v1/gonum/mat 우수한 행렬 구성기능 패키지 사용한 점

func main() {

	// 행렬의 수평적인 표현을 생성한다
	data := []float64{1.2, -5.7, -2.4, 7.3, 4.0, 5.0, 3.0, 2.0, 1.0}

	// 행렬을 만든다(첫 번째 인자는 행의 수이며, 두 번째 인자는
	// 열의 수를 나타낸다)
	a := mat64.NewDense(3, 3, data)

	// 검사를 위해 행렬을 표준 출력을 통해 출력한다.
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n\n", fa)
}

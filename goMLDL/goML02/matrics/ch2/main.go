package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// 행렬의 수평적인 표현을 생성한다.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// 행렬을 만든다(첫 번째 인자는 행의 수, 두 번째 인자는 열의 수를 나타낸다)
	a := mat64.NewDense(2, 2, data)

	// 행렬에서 값 하나를 가져온다.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)

	// 특정 열에서 값을 가져온다.
	col := mat64.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	// 특정 행에서 값을 가져온다.
	row := mat64.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	// 행렬의 값 하나를 변경한다.
	a.Set(0, 1, 11.2)

	// 전체 행을 변경한다.
	a.SetRow(0, []float64{14.3, -4.2})

	// 전체 열을 변경한다.
	a.SetCol(0, []float64{1.7, -0.3})

	// 검사를 위해 행렬을 표준 출력을 통해 출력한다.
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
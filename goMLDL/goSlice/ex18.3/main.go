//ch18/ex18.3/ex18.3.go
package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 10; i++ { // ❶ 요소를 하나씩 추가
		slice = append(slice, i)  // 이단계 Slice는 1부터 10의 요소값을 갖는
	}

	slice = append(slice, 11, 12, 13, 14, 15) // Slice 1-15값는 새로운슬라이스 ❷ 한 번에 여러 요소 추가
	fmt.Println(slice)
}

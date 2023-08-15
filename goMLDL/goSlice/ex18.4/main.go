//ch18/ex18.4/ex18.4.go
package main

import "fmt"

func changeArray(array2 [5]int) { // ❶ 배열을 받아서 인덱스 세 번째 값 변경
	array2[2] = 200 // C/C++ 배열에서는 인덱스3번째 요소값이 변경된다.하지만 golan 배열에서는 변경되지 않는다.
}

func changeSlice(slice2 []int) { // ❷ go슬라이스를 받아서 인덱스 세 번째 값 변경
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}

package main

import "fmt"

func main() {
    //동적배열 슬라이스에 초기값을 넣어주어 에러를 없앴다
	var slice = []int{1, 2, 3} // ❶ 동적배열 초기값 요소가 3개인 슬라이스
    //slice := make([]int, 3)

	slice2 := append(slice, 4) // ❷ 요소 추가

	fmt.Println(slice)
	fmt.Println(slice2)
}

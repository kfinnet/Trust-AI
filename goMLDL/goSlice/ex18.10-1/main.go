package main

import (
	"fmt"
)

func addNum(slice *[]int) {
	*slice = append(*slice, 4)  // * 포인터 처리가 답, cap(4)
}

func main() {
	slice := []int{1, 2, 3} // 포인터가 len(3) cap(3) 공간이 없다.
	addNum(&slice) //&Slice 포인터로cap(4)추가된 새로운 슬라이스

	fmt.Println(slice)  
}
package main

import "fmt"

func main() {

	// 슬라이스(slice)를 통해 "벡터"를 초기화한다.
	var myvector []float64

	// 벡터에 구성요소를 몇 개 추가한다.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	fmt.Println(myvector)
}

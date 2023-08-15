package main

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

func main() {

	// 기본 만료 시간이 5분이며 30초마다 만료된 항목을 제거하는
	// 캐시를 생성한다
	c := cache.New(5*time.Minute, 30*time.Second)

	// 캐시에 키-값(Key-Value)을 넣는다
	c.Set("mykey", "myvalue", cache.DefaultExpiration)

	// 캐시에서 mykey 값을 가져올 때는 Get 메소드를 사용하면 된다
	v, found := c.Get("mykey")
	if found {
		fmt.Printf("key: mykey, value: %s\n", v)
	}
}

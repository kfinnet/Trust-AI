package main

import (
	"fmt"
	"log"
	//cache "github.com/patrickmn/go-cache"
	"github.com/boltdb/bolt"
)

func main() {

	// 현재 디렉토리에서 embedded.db 데이터 파일을 연다.
	// 파일이 존재하지 않는 경우에는 파일을 생성한다.
	db, err := bolt.Open("embedded.db", 0777, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 데이터를 저장하기 위해 boltdb 파일에 "bucket"을 생성한다
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// BoltDB 파일에 키-값 조합의 데이터를 넣는다
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// BoltDB 파일에 저장된 키-값 조합의 데이터를 
	// 표준출력으로 출력한다
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

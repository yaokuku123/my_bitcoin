package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// 1.打开 boltdb
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	// 2.关闭 boltdb
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		// 3.找到 bucket
		bucket := tx.Bucket([]byte("demo_db"))
		// 未找到 bucket 则创建
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("demo_db"))
			if err != nil {
				log.Panic(err)
			}
		}
		// 4.写数据
		bucket.Put([]byte("key1"), []byte("value1"))
		bucket.Put([]byte("key2"), []byte("value2"))
		return nil
	})
	// 5.读数据
	db.View(func(tx *bolt.Tx) error {
		// 找到 bucket
		bucket := tx.Bucket([]byte("demo_db"))
		if bucket == nil {
			log.Panic("未找到bucket")
		}
		res1 := bucket.Get([]byte("key1"))
		res2 := bucket.Get([]byte("key2"))
		fmt.Printf("res1: %s\n", res1)
		fmt.Printf("res2: %s\n", res2)
		return nil
	})
}

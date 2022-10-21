package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Dog struct {
	Name string
}

type Human struct {
	Gender string
}

type Person struct {
	Human
	Name  string
	Age   int
	MyDog Dog
}

func main() {
	p := Person{Human: Human{"男"}, Name: "yorick", Age: 25, MyDog: Dog{"tony"}}
	// 编码
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&p)
	if err != nil {
		log.Panic(err)
	}
	// 解码
	var decodeP Person
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	decoder.Decode(&decodeP)
	fmt.Println(decodeP)
}

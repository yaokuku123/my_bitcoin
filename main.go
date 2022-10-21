package main

import (
	"fmt"
	"my_bitcoin/blockchain"
)

func main() {
	block := blockchain.NewBlock("hello world", []byte{})
	fmt.Printf("prevHash:%v\n", block.PrevHash)
	fmt.Printf("hash:%v\n", block.Hash)
	fmt.Printf("data:%s\n", block.Data)
}

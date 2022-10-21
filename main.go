package main

import (
	"fmt"
	"my_bitcoin/blockchain"
)

func main() {
	block := blockchain.NewBlock("hello world", []byte{})
	fmt.Println(block.PrevHash)
	fmt.Println(block.Hash)
	fmt.Println(block.Data)
}

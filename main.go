package main

import (
	"fmt"
	"my_bitcoin/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	for _, block := range bc.Blocks {
		fmt.Printf("prevHash:%v\n", block.PrevHash)
		fmt.Printf("hash:%v\n", block.Hash)
		fmt.Printf("data:%s\n", block.Data)
	}
}

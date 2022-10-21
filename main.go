package main

import (
	"fmt"
	"my_bitcoin/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("transfer 50 coin")
	bc.AddBlock("transfer 20 coin")
	for i, block := range bc.Blocks {
		fmt.Printf("----------%d-----------\n", i)
		fmt.Printf("prevHash:%x\n", block.PrevHash)
		fmt.Printf("hash:%x\n", block.Hash)
		fmt.Printf("data:%s\n", block.Data)
	}
}

package main

import (
	"my_bitcoin/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("transfer 50 coin")
	bc.AddBlock("transfer 20 coin")
}

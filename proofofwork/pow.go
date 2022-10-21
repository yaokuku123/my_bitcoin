package proofofwork

import (
	"math/big"
	"my_bitcoin/blockchain"
)

type ProofOfWork struct {
	Block  *blockchain.Block // 区块
	Target *big.Int          // 挖矿目标值
}

// NewProofOfWork 创建工作量证明，主要工作是放入block并设置挖矿目标值
func NewProofOfWork(block *blockchain.Block) *ProofOfWork {
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	target := big.Int{}
	target.SetString(targetStr, 16)
	pow := ProofOfWork{
		Block:  block,
		Target: &target,
	}
	return &pow
}

// Run 运行pow挖矿，获取可以打包的hash值
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// TODO
	return []byte{}, 0
}

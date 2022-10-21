package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"my_bitcoin/util"
)

type ProofOfWork struct {
	Block  *Block   // 区块
	Target *big.Int // 挖矿目标值
}

// NewProofOfWork 创建工作量证明，主要工作是放入block并设置挖矿目标值
func NewProofOfWork(block *Block) *ProofOfWork {
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
	var nonce uint64 = 0
	bc := pow.Block
	var hash [32]byte
	// 1.循环nonce
	for {
		// 2.拼装被hash的数据
		blockInfoList := [][]byte{
			util.Uint64ToByte(bc.Version),
			bc.PrevHash,
			bc.MerkleRoot,
			util.Uint64ToByte(bc.TimeStamp),
			util.Uint64ToByte(bc.Difficulty),
			util.Uint64ToByte(nonce),
			bc.Data,
		}
		blockInfo := bytes.Join(blockInfoList, []byte{})
		// 3.sha256
		hash = sha256.Sum256(blockInfo)
		// 4.与目标hash进行比较，若小于目标值则说明找到hash；否则，继续遍历随机值计算下一个hash
		tmpValue := big.Int{}
		tmpValue.SetBytes(hash[:])
		if tmpValue.Cmp(pow.Target) == -1 {
			//fmt.Printf("挖矿成功，nonce：%d\n", nonce)
			break
		}
		nonce++
	}
	return hash[:], nonce
}

package blockchain

import (
	"bytes"
	"crypto/sha256"
	"my_bitcoin/proofofwork"
	"my_bitcoin/util"
	"time"
)

type Block struct {
	Version    uint64 //版本号
	PrevHash   []byte //前区块hash
	MerkleRoot []byte // merkle根
	TimeStamp  uint64 //时间戳
	Difficulty uint64 //难度值
	Nonce      uint64 //随机数

	Hash []byte //hash
	Data []byte //数据
}

// NewBlock 新建区块
func NewBlock(data string, prevHash []byte) *Block {
	// 1.构建新区块
	block := Block{
		Version:    00,
		PrevHash:   prevHash,
		MerkleRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	// 2.将区块传入到共识算法对象中
	pow := proofofwork.NewProofOfWork(&block)
	// 3.挖矿，得到最后符合要求的 hash 和 随机数
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

// GenesisBlock 创世纪区块
func GenesisBlock() *Block {
	genesisBlock := NewBlock("first block", []byte{})
	return genesisBlock
}

// SetHash 设置 Hash
func (bc *Block) SetHash() {
	// 拼接需要被 Hash 的数据
	blockInfoList := [][]byte{
		util.Uint64ToByte(bc.Version),
		bc.PrevHash,
		bc.MerkleRoot,
		util.Uint64ToByte(bc.TimeStamp),
		util.Uint64ToByte(bc.Difficulty),
		util.Uint64ToByte(bc.Nonce),
		bc.Data,
	}
	blockInfo := bytes.Join(blockInfoList, []byte{})
	// sha256
	hash := sha256.Sum256(blockInfo)
	// 将 hash 设置到 block 中
	bc.Hash = hash[:]
}

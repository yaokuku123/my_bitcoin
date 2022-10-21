package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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
	pow := NewProofOfWork(&block)
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
func (block *Block) SetHash() {
	// 拼接需要被 Hash 的数据
	blockInfoList := [][]byte{
		util.Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkleRoot,
		util.Uint64ToByte(block.TimeStamp),
		util.Uint64ToByte(block.Difficulty),
		util.Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(blockInfoList, []byte{})
	// sha256
	hash := sha256.Sum256(blockInfo)
	// 将 hash 设置到 block 中
	block.Hash = hash[:]
}

// Serialize 序列化区块
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// Deserialize 反序列化区块
func Deserialize(data []byte) Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return block
}

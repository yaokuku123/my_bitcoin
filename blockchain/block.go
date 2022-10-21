package blockchain

import "crypto/sha256"

type Block struct {
	PrevHash []byte //前区块hash
	Hash     []byte //hash
	Data     []byte //数据
}

// NewBlock 新建区块
func NewBlock(data string, prevHash []byte) *Block {
	block := Block{
		PrevHash: prevHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	block.SetHash()
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
	data := append(bc.PrevHash, bc.Data...)
	// sha256
	hash := sha256.Sum256(data)
	// 将 hash 设置到 block 中
	bc.Hash = hash[:]
}

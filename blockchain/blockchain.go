package blockchain

type BlockChain struct {
	Blocks []*Block // 区块数组
}

// NewBlockChain 创建区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	blockChain := BlockChain{
		Blocks: []*Block{genesisBlock},
	}
	return &blockChain
}

// AddBlock 添加新区块到区块链
func (bc *BlockChain) AddBlock(data string) {
	// 1.获取前区块的 hash
	prevHash := bc.Blocks[len(bc.Blocks)-1].Hash
	// 2.创建新区块
	block := NewBlock(data, prevHash)
	// 3.将新区块添加到区块链中
	bc.Blocks = append(bc.Blocks, block)
}

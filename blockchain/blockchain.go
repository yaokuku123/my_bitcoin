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

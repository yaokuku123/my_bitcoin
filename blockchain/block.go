package blockchain

type Block struct {
	PrevHash []byte //前区块hash
	Hash     []byte //hash
	Data     []byte //数据
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{
		PrevHash: prevHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	return &block
}

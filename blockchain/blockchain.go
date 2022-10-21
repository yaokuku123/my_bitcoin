package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

const DbPath = "blockChain.db"
const BucketName = "blockChain_data"
const LastHash = "lastHash"

type BlockChain struct {
	DB   *bolt.DB // 数据库
	tail []byte   // 最后区块 hash
}

// NewBlockChain 创建区块链
func NewBlockChain() *BlockChain {
	var lasthash []byte
	// 从数据库中读取区块信息
	// 1.打开 boltdb 数据库
	db, err := bolt.Open(DbPath, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		// 2.找到 bucket
		bucket := tx.Bucket([]byte(BucketName))
		// 3.未找到 bucket 则创建
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BucketName))
			if err != nil {
				log.Panic(err)
			}
			// 3.1 创建创世纪区块
			genesisBlock := GenesisBlock()
			// 3.2 将创世纪区块序列化后放入数据库
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			// 3.3 最后一个区块的 hash 也需要放入数据库
			bucket.Put([]byte(LastHash), genesisBlock.Hash)
			// 内存中也放一份
			lasthash = genesisBlock.Hash
		} else {
			// 找到 bucket 说明不是第一次创建区块链，则直接加载当前系统的最后一个区块 hash
			lasthash = bucket.Get([]byte(LastHash))
		}
		return nil
	})
	return &BlockChain{db, lasthash}
}

// AddBlock 添加新区块到区块链
func (bc *BlockChain) AddBlock(data string) {
	// 1.获取前区块的 hash
	db := bc.DB
	prevHash := bc.tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket == nil {
			log.Panic("bucket 获取失败")
		}
		// 2.创建新区块
		block := NewBlock(data, prevHash)
		// 3.将新区块和最后区块 hash 添加到区块链中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(LastHash), block.Hash)
		// 更新内存的最后区块 hash
		bc.tail = block.Hash
		return nil
	})

}

package blockchain

import (

	"github.com/boltdb/bolt"

)

const LAST_HASH = "lasthash"
const BUCKET_NAME = "blocks"
const BLOCKCHAIN = "blockchain.db"

/**
 *区块链结构体的定义，代表的是一条区块链
 */
//1、将新区块数据与已有区块进行连接
//2、查询某个区块的数据和信息
//3、遍历区块信息


type Blockchain struct {
	LastHash []byte  //表示区块链中最新区块的哈希，用于查找最新的区块内容
	BoltDb *bolt.DB  //区块链中操作区块数据文件的数据库操作对象
}

/**
 *创建一条区块链
 */
func NewBlockchain() Blockchain{
	//创世区块
	genesis := CreateGenesisBlock()
	db, err := bolt.Open("blockchain.db",0600,nil)
	if err != nil {
		panic(err.Error())
	}
	bc := Blockchain{
		LastHash: genesis.Hash,
		BoltDb:   db,
	}
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(BUCKET_NAME))
		if err != nil {
			panic(err.Error())
		}
		//序列化
		genesisBytes := genesis.Serialize()
		//把创世区块存储到桶中
		bucket.Put(genesis.Hash, genesisBytes)
		bucket.Put([]byte(LAST_HASH), genesis.Hash)
		//更新最新区块的
		return nil
	})
	return bc
}


/**
 *保存数据到区块链中：先生产一个新区块，然后将新区块添加到区块中
 */
func (bc  Blockchain) SaveData(data []byte) {
	//1、从文件中读取到最新区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("读取区块链数据失败")
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(lastHash)
		//反序列化

		lastBlock, _ := DeSerialize(lastBlockBytes)
		return nil
	})
	//2、新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块存到文件中
	db.Update()
}

package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"time"
)

/**
* 定义区块结构体，用于表示区块
 */
type Block struct {
	Height int64 //区块的高度，第几个区块
	TimeStamp int64 //区块链时间戳
	PrevHash []byte //前一个区块的哈希
	Data []byte //数据字段
	Hash []byte //当前区块的哈希值
	Version string //版本号
	Nonce int64 //区块对应的nonce值
}

/**
*创建一个新区块
 */
func NewBlock(height int64, prevHash []byte, data []byte) Block{
		block := Block {
			Height: height,
			TimeStamp: time.Now().Unix(),
			PrevHash: prevHash,
			Data: data,
			Version: "0x01",
		}

		pow := NewPow(block)
		hash,nonce := pow.Run()
		block.Nonce = nonce
		fmt.Println("挖矿挖矿到的hash:",hash)
		block.Hash = hash

	//1.将block结构体数据类型转换为[]byte类型
	//heightBytes,_ := Int64ToByte(block.Height)
	//timeStampBytes,_ := Int64ToByte(block.TimeStamp)
	//versionBytes := StringToBytes(block.Version)
	//nonceBytes,_ := utils.Int64ToByte(block.Nonce)
	//
	//var blockBytes []byte
	////bytes.Join 拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	//}, []byte{})
	//
	//	//调用Hash计算，对区块进行sha256哈希计算
	//	block.Hash = utils.SHA256HashBlock(blockBytes)

		//挖矿竞争，获得记帐权
		return block
}
 /**
 *创建创世区块
  */
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, nil)
	return genesisBlock
}


/**
 * 对区块进行序列化操作
 */
func (b Block) Serialize() ([]byte){
	buff := new(bytes.Buffer)  //缓冲区
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b)  //将区块b放入到序列化编码器中
	return buff.Bytes()
}

/**
 * 区块反序列化操作
 */
func DeSerialize(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil,err
	}
	return &block, nil
}


func Int64ToByte(num int64) ([]byte, error) {
	//Buffer:  缓冲区
	buff := new(bytes.Buffer)  //通过new实例化一个缓冲区
	//buff.Write()  通过一系列的write方法向缓冲区写入数据
	//buff.Bytes()   通过bytes方法从缓冲区中获取数据
	/**
	*两种排序方式：
	*  大端位序排列：BigEndian
	*  小端位序排列：LittleEndian
	 */
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	//从缓冲区读取数据
	return buff.Bytes(), nil
}

/**
 *将字符串转换为[]byte
 */

func StringToBytes(data string) []byte {
	return []byte(data)
}

//1.计算区块哈希
//2.挖矿

package main

import (
	"BCP/blockchain"
	"BCP/db_mysql"
	_ "BCP/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)


func main() {

	block0 := blockchain.CreateGenesisBlock()  //创建创世区块
	//block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	fmt.Println(block0)
	fmt.Printf("block0的哈希:%x\n", block0.Hash )
	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block1的哈希:%x\n",block1.Hash)
	fmt.Printf("block1的PrevHash:%x\n",block1.PrevHash)


	block0Bytes := block0.Serialize()
	fmt.Println("创世区块gob序列化后:" ,block0Bytes)
	deBlock0, err := blockchain.DeSerialize(block0Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("反序列化后的区块高度是",deBlock0.Height)
	//序列化
	blockJson, _ := json.Marshal(block0)
	fmt.Println("序通过JSON列化以后的block:",string(blockJson))

	blockXml, _:= xml.Marshal(block0)
	fmt.Println("通过xml序列化以后的block:", string(blockXml))
	return


	db_mysql.Connect()

	beego.SetStaticPath("/js","./static.js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()



}


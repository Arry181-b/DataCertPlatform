package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/**
*对一个字符串数据进行MD5哈希运算
 */
func MD5HashString(data string) string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

func MD5HashReader(reader io.Reader) (string,error) {
	md5Hash := md5.New()
	readerBytes, err := ioutil.ReadAll(reader)
	//fmt.Println("读取到的文件",readerBytes)
	if err != nil {
		return "", err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

/**
 * 读取io流中的数据，并对数据进行哈希计算，返回sha256哈希值
 */
func SHA256HashReader (reader io.Reader) (string, error){
	sha256Hash := sha256.New()
	readerBytes, err := ioutil.ReadAll(reader)
	fmt.Println("读取到的文件：", readerBytes)
	if err != nil {
		return"", err
	}
	sha256Hash.Write(readerBytes)
	hashBytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}


/**
 * 对区块数据进行SHA256Hash计算
 */
func SHA256HashBlock(bs []byte) []byte {

	//2.将转换后的[]byte字节切片输入write方法
	sha256Hash := sha256.New()
	sha256Hash.Write(bs)
	hash := sha256Hash.Sum(nil)
	return  hash
}















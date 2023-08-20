package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Sign(content string) string {
	message := []byte(content) //字符串转化字节数组
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New() //sha-256加密
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

func SHA256Verify(content, sign string) bool {
	return SHA256Sign(content) == sign
}

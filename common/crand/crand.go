package crand

import (
	"bytes"
	"crypto/rand"
	"math/big"
)

const (
	baseLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	baseNum     = "1234567890"
)

//Letters 随机固定长度字符串
func Letters(n int) string {
	var container string
	b := bytes.NewBufferString(baseLetters)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < n; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(baseLetters[randomInt.Int64()])
	}
	return container
}
//Num 随机固定长度数字组合
func Num(n int) string {
	var container string
	b := bytes.NewBufferString(baseNum)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < n; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(baseNum[randomInt.Int64()])
	}
	return container
}

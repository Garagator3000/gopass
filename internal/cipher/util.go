package cipher

import (
	"bytes"
	"crypto/sha256"
)

func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func unpad(data []byte) []byte {
	padding := data[len(data)-1]
	return data[:len(data)-int(padding)]
}

func hashKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

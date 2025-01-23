package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func Encrypt(key string, data string) string {
	keyHash := hashKey(key)

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return ""
	}

	paddedData := pad([]byte(data), block.BlockSize())

	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)

	return hex.EncodeToString(ciphertext)
}

func Decrypt(key string, data string) string {
	keyHash := hashKey(key)

	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return ""
	}

	if len(ciphertext) < aes.BlockSize {
		return ""
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(unpad(ciphertext))
}

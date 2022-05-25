package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func AESEncrypt(context string, ik, iv []byte) (string, error) {
	s := []byte(context)
	block, err := aes.NewCipher(ik)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	s = PKCS7Padding(s, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(s))
	blockMode.CryptBlocks(crypted, s)
	return hex.EncodeToString(crypted), nil
}

func AESDecrypt(content string, ik, iv []byte) (string, error) {
	s, _ := hex.DecodeString(content)
	block, err := aes.NewCipher(ik)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	origData := make([]byte, len(s))
	blockMode.CryptBlocks(origData, s)
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

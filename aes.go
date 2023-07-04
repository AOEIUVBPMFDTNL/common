package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func EncryptKey(key []byte, b ...byte) []byte {
	h := md5.New()
	h.Write(key)
	r := md5.New()
	p := hex.EncodeToString(h.Sum(b))
	r.Write([]byte(p[12:32]))
	return []byte(hex.EncodeToString(r.Sum(b)))
}

func AesEncrypt(data, key []byte) (result string, err error) {
	key = EncryptKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	data = PKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	result = base64.StdEncoding.EncodeToString(crypted)
	return
}

func AesDecrypt(data string, key []byte) (result []byte, err error) {
	result, err = base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	key = EncryptKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(result))
	blockMode.CryptBlocks(origData, result)
	result = PKCS7UnPadding(origData)
	return
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

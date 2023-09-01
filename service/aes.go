package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
)

type Aes struct {
	Key string
}

// Sha256Key sha256 加密
func (a *Aes) Sha256Key(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	newKey := h.Sum(nil)
	return newKey
}

// PKCS7Padding 填充数据
func (a *Aes) PKCS7Padding(ciphertext []byte) []byte {
	bs := aes.BlockSize
	padding := bs - len(ciphertext)%bs
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

// PKCS7UnPadding 放出数据
func (a *Aes) PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密
func (a *Aes) Encrypt(origData string) (string, error) {
	newKey := a.Sha256Key(a.Key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	newOrigData := []byte(origData)
	newOrigData = a.PKCS7Padding(newOrigData)
	blockMode := cipher.NewCBCEncrypter(block, newKey[:16])
	crypted := make([]byte, len(newOrigData))
	blockMode.CryptBlocks(crypted, newOrigData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt 解密
func (a *Aes) Decrypt(crypted string) (string, error) {
	newKey := a.Sha256Key(a.Key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	newCrypted, _ := base64.StdEncoding.DecodeString(crypted)
	blockMode := cipher.NewCBCDecrypter(block, newKey[:16])
	origData := make([]byte, len(newCrypted))
	blockMode.CryptBlocks(origData, newCrypted)
	origData = a.PKCS7UnPadding(origData)
	return string(origData), nil
}

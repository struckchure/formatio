package lib

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

type ecbEncrypter ecb

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ThressDSEncrypter struct{}

func (t *ThressDSEncrypter) Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func (t *ThressDSEncrypter) EncryptData(key, plainText string) (string, error) {
	keyBytes := []byte(key)
	plainTextBytes := []byte(plainText)

	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		return "", err
	}

	paddedText := t.Pkcs5Padding(plainTextBytes, block.BlockSize())

	cipherText := make([]byte, len(paddedText))
	encrypter := NewECBEncrypter(block)
	encrypter.CryptBlocks(cipherText, paddedText)

	encoded := base64.StdEncoding.EncodeToString(cipherText)
	return encoded, nil
}

// Example
//
//	encrypter := NewThressDSEncrypter()
//	encrypted, err := encrypter.EncryptData(key, plainText)
//	fmt.Println(encrypted)
func NewThressDSEncrypter() ThressDSEncrypter {
	return ThressDSEncrypter{}
}

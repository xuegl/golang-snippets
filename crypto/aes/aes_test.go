package aes

import (
	aes2 "crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	key := []byte("12345678123456781234567812345678")
	aesCipher, err := aes2.NewCipher(key)
	if err != nil {
		t.Error(err)
	}
	iv := []byte("abcdefghijklmnop")
	cbcEncrypter := cipher.NewCBCEncrypter(aesCipher, iv)

	plainText := []byte("I love go programming language!!")
	cipherTextWithIV := make([]byte, aes2.BlockSize+len(plainText))
	cbcEncrypter.CryptBlocks(cipherTextWithIV[aes2.BlockSize:], plainText)
	copy(cipherTextWithIV[:aes2.BlockSize], iv)
	fmt.Printf("plain text: %s\n", plainText)
	fmt.Printf("cipher text (with iv): %x\n", cipherTextWithIV)
	// output: 6162636465666768696a6b6c6d6e6f70bc93b5cb1a081b47357f73d40966e3ce53c29db21a13bec2f9be4f76d8f09f2b
}

func TestAesDecrypt(t *testing.T) {
	key := []byte("12345678123456781234567812345678")
	aesCipher, err := aes2.NewCipher(key)
	if err != nil {
		t.Error(err)
	}
	cipherTextWithIV, err := hex.DecodeString("6162636465666768696a6b6c6d6e6f70bc93b5cb1a081b47357f73d40966e3ce53c29db21a13bec2f9be4f76d8f09f2b")
	if err != nil {
		t.Error(err)
	}

	iv := cipherTextWithIV[:aes2.BlockSize]
	cipherText := cipherTextWithIV[aes2.BlockSize:]
	plainText := make([]byte, len(cipherText))
	cbcDecrypter := cipher.NewCBCDecrypter(aesCipher, iv)
	cbcDecrypter.CryptBlocks(plainText, cipherText)
	fmt.Printf("cipher text (with iv): %s\n", cipherTextWithIV)
	fmt.Printf("plain text: %s\n", plainText)
}

func TestAESCipher_Encrypt(t *testing.T) {
	key := []byte("12345678123456781234567812345678")
	iv := []byte("abcdefghijklmnop")
	aesCipher, err := NewAESCipher(key, iv)
	if err != nil {
		t.Error(err)
	}
	originalText := []byte("Hello World")
	cipherText := aesCipher.Encrypt(originalText)
	plainText, err := aesCipher.Decrypt(cipherText)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, originalText, plainText)
}

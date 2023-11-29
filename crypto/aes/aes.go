package aes

import (
	"bytes"
	aes2 "crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type AESCipher struct {
	key       []byte
	iv        []byte
	aesCipher cipher.Block
}

func NewAESCipher(key []byte, iv []byte) (*AESCipher, error) {
	switch l := len(key); l {
	default:
		return nil, aes2.KeySizeError(l)
	case 16, 24, 32:
	}
	switch l := len(iv); l {
	default:
		return nil, errors.New("iv length must be equal to block size")
	case aes2.BlockSize:
	}

	aesCipher, err := aes2.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESCipher{
		key:       key,
		iv:        iv,
		aesCipher: aesCipher,
	}, nil
}

func (a *AESCipher) Encrypt(src []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	cbcEncryptor := cipher.NewCBCEncrypter(a.aesCipher, a.iv)
	if len(src)%aes2.BlockSize == 0 {
		dst := make([]byte, len(src))
		cbcEncryptor.CryptBlocks(dst, src)
		return dst
	} else {
		src = a.pkcs7Pad(src, aes2.BlockSize)
		dst := make([]byte, len(src))
		cbcEncryptor.CryptBlocks(dst, src)
		return dst
	}
}

func (a *AESCipher) Decrypt(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return nil, nil
	}
	cbcDecrypter := cipher.NewCBCDecrypter(a.aesCipher, a.iv)
	if len(src)%aes2.BlockSize != 0 {
		return nil, fmt.Errorf("input is not full block size")
	}
	dst := make([]byte, len(src))
	cbcDecrypter.CryptBlocks(dst, src)
	return a.pkcs7Unpad(dst), nil
}

func (a *AESCipher) pkcs7Pad(src []byte, size int) []byte {
	i := aes2.BlockSize - (len(src) % aes2.BlockSize)
	return append(src, bytes.Repeat([]byte{byte(i)}, i)...)
}

func (a *AESCipher) pkcs7Unpad(src []byte) []byte {
	i := int(src[len(src)-1]) % aes2.BlockSize
	return src[:len(src)-i]
}

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type AESCipher struct {
	key       []byte
	iv        []byte
	aesCipher cipher.Block
}

func NewAESCipher(key []byte, iv []byte) (*AESCipher, error) {
	switch l := len(key); l {
	default:
		return nil, aes.KeySizeError(l)
	case 16, 24, 32:
	}
	switch l := len(iv); l {
	default:
		return nil, errors.New("iv length must be equal to block size")
	case aes.BlockSize:
	}

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AESCipher{
		key:       key,
		iv:        iv,
		aesCipher: aesCipher,
	}, nil
}

func (a *AESCipher) Encrypt(src []byte) {
	//cbcEncryptor := cipher.NewCBCEncrypter(a.aesCipher, a.iv)
	for i := 0; i < len(src); i += aes.BlockSize {

	}
}

func (a *AESCipher) Decrypt() {
	//cbcDecrypter := cipher.NewCBCDecrypter(a.aesCipher, a.iv)

}

func (a *AESCipher) pkcs7Pad() {

}

func (a *AESCipher) pkcs7Unpad() {

}

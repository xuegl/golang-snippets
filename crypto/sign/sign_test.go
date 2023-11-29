package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err)
	}
	publicKey := privateKey.PublicKey

	msg := []byte("I love go programming language!!")
	dgst := sha256.Sum256(msg)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dgst[:])
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("sign: %s\n", hex.EncodeToString(sign))

	err = rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, dgst[:], sign)
	if err != nil {
		t.Error(err)
	}
}

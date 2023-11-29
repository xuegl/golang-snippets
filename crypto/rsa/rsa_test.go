package rsa

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	privateKey, err := GenerateKey(2048)
	if err != nil {
		t.Error(err)
	}
	publicKey := privateKey.PublicKey
	fmt.Printf("Private Key's size = %d bits\n", privateKey.Size()*8) // 2048
	fmt.Printf("Public Key's size = %d bits\n", publicKey.Size()*8)   // 2048”
}

package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHmac(t *testing.T) {
	key := []byte("12345678123456781234567812345678")
	msg := []byte("I love go programming language!!")
	m := hmac.New(sha256.New, key)
	m.Write(msg)
	sum := m.Sum(nil)
	fmt.Printf("mac: %s\n", hex.EncodeToString(sum))
}

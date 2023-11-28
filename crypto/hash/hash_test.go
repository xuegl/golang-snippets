package hash

import (
	"crypto"
	"encoding/hex"
	"testing"
)

func TestHash(t *testing.T) {
	ret, err := Hash([]byte{'H', 'e', 'l', 'l', 'o'}, crypto.MD5)
	if err != nil {
		t.Error(err)
	}
	println(hex.EncodeToString(ret))
}

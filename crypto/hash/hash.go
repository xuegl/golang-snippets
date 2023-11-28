package hash

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"fmt"
)

func Hash(src []byte, hash crypto.Hash) ([]byte, error) {
	if !hash.Available() {
		return nil, fmt.Errorf("%s is not available", hash.String())
	}
	h := crypto.Hash.New(hash)
	_, err := h.Write(src)
	if err != nil {
		return nil, err
	}
	ret := h.Sum(nil)
	return ret, nil
}

package strings

import (
	"crypto/rand"
)

const (
	RandomSymbols       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomSymbolsLength = byte(len(RandomSymbols))
)

// Rand produce a random string with the given length l
func Rand(l uint) (string, error) {
	buf := make([]byte, l)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}

	for i, b := range buf {
		buf[i] = RandomSymbols[b%RandomSymbolsLength]
	}

	return string(buf), nil
}

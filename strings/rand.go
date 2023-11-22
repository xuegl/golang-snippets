package strings

import "crypto/rand"

var (
	RandomSymbols       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RandomSymbolsLength = byte(len(RandomSymbols))
)

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

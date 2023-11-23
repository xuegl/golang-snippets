package tests

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestBig(t *testing.T) {
	id := time.Now().UnixMicro()
	s := big.NewInt(id).Text(62)
	fmt.Println(s) // 18Ik

	id2 := big.Int{}
	id2.SetString(s, 62)
	fmt.Println(id2) // 271828
}

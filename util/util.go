package util

import (
	"fmt"
	"time"
)

func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)
}

func setupTeardown() func() {
	fmt.Println("Run initialization")
	return func() {
		fmt.Println("Run cleanup")
	}
}

func Int64toB(i int64) []byte {
	buf := make([]byte, 8, 8)
	for k := 0; k < 8; k++ {
		buf[k] = byte(i & 0xFF)
		i = i >> 8
	}

	return buf
}

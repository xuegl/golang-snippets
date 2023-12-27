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

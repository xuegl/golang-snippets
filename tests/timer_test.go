package tests

import (
	"testing"
	"time"
)

func TestTimer_Stop(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	time.Sleep(2 * time.Second)

	if !timer.Stop() {
		<-timer.C
		println("timer expired")
	}
}

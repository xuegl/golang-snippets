package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestBreakLabel(t *testing.T) {
	exit := make(chan struct{})
	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				break loop
			}
		}
		fmt.Println("exit!")
	}()

	time.Sleep(3 * time.Second)
	exit <- struct{}{}

	time.Sleep(3 * time.Second)
}

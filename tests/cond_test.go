package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	type signal struct{}
	var ready bool
	worker := func(i int) {
		fmt.Printf("worker %d is running...\n", i)
		time.Sleep(time.Second)
		fmt.Printf("worker %d completed\n", i)
	}

	spawnGroup := func(f func(int), cnt int, cond *sync.Cond) <-chan signal {
		c := make(chan signal)
		var wg sync.WaitGroup

		for i := 0; i < cnt; i++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				cond.L.Lock()
				for !ready {
					cond.Wait()
				}
				cond.L.Unlock()
				fmt.Printf("worker %d started to work...\n", j)
				f(j)
			}(i)
		}

		go func() {
			wg.Wait()
			c <- signal{}
		}()
		return c
	}

	fmt.Println("start a group of workers...")
	cond := sync.NewCond(&sync.Mutex{})
	c := spawnGroup(worker, 5, cond)

	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	cond.L.Lock()
	ready = true
	cond.Broadcast()
	cond.L.Unlock()

	<-c
	fmt.Println("the group of workers completed")
}

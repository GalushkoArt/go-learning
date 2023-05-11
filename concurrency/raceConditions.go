package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func RaceConditions() {
	waitGroup := sync.WaitGroup{}
	var counter uint32
	start := time.Now()
	for i := 0; i < 100000; i++ {
		waitGroup.Add(1)
		go func() {
			counter++
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	end := time.Now()
	fmt.Println("After not synchronised increment", counter, "and it took", end.Sub(start))
	counter = 0
	mu := new(sync.RWMutex)
	start = time.Now()
	for i := 0; i < 100000; i++ {
		waitGroup.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	end = time.Now()
	mu.RLock()
	fmt.Println("After increment with mutex", counter, "and it took", end.Sub(start))
	mu.RUnlock()

	counter = 0
	start = time.Now()
	for i := 0; i < 100000; i++ {
		waitGroup.Add(1)
		go func() {
			atomic.AddUint32(&counter, 1)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	end = time.Now()
	fmt.Println("After increment with atomic", atomic.LoadUint32(&counter), "and it took", end.Sub(start))
}

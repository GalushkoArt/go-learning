package concurrency

import (
	"fmt"
	"time"
)

func BasicGoroutine() {
	go fmt.Println("Goroutine 1 finished")
	go fmt.Println("Goroutine 2 finished")
	go fmt.Println("Goroutine 3 finished")
	go fmt.Println("Goroutine 4 finished")

	time.Sleep(time.Millisecond)

	fmt.Println("Main is here!")
}

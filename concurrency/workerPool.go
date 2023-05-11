package concurrency

import (
	"fmt"
	"time"
)

func WorkerPool() {
	start := time.Now()
	const numberOfWorkers, numberOfJobs = 5, 15

	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	for i := 0; i < numberOfWorkers; i++ {
		go worker(i+1, jobs, results)
	}

	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numberOfJobs; i++ {
		fmt.Printf("Result #%d is %d\n", i, <-results)
	}
	fmt.Printf("Time Elapsed for %d jobs with %d workers: %s\n", numberOfJobs, numberOfWorkers, time.Since(start))
}

func worker(workerId int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d has finished job\n", workerId)
		result <- j * j
	}
}

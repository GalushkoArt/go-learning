package tools

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

func MakeLoad(requestNumber int64, workerNumber int, timeout time.Duration, url string, method string, body io.Reader) error {
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	tasks := make(chan func(), workerNumber)
	timeTrack := make(chan int64, 2*workerNumber)
	timeStorage := make([]int64, 0, requestNumber)
	mu := sync.Mutex{}
	for i := 0; i < workerNumber; i++ {
		go func() {
			for task := range tasks {
				task()
			}
		}()
	}
	for i := 0; i < 20; i++ {
		go func() {
			for elapsed := range timeTrack {
				mu.Lock()
				timeStorage = append(timeStorage, elapsed)
				mu.Unlock()
			}
		}()
	}
	wg := sync.WaitGroup{}
	var errorNumber int64
	log.Println("Start load")
	start := time.Now()
	for i := int64(0); i < requestNumber; i++ {
		wg.Add(1)
		tasks <- func() {
			start := time.Now()
			do, err := client.Do(request)
			elapsed := time.Since(start).Nanoseconds()
			timeTrack <- elapsed
			if err != nil {
				atomic.AddInt64(&errorNumber, 1)
				log.Println("Found error!", err)
				wg.Done()
				return
			}
			if do.StatusCode != 200 {
				atomic.AddInt64(&errorNumber, 1)
				log.Println("Wrong status code!", do.StatusCode)
				wg.Done()
				return
			}
			body, err := io.ReadAll(do.Body)
			if err != nil || len(body) == 0 {
				atomic.AddInt64(&errorNumber, 1)
				log.Println("Error in response body!", err, string(body))
				wg.Done()
				return
			}
			wg.Done()
		}
	}
	wg.Wait()
	elapsed := time.Since(start)
	close(tasks)
	time.Sleep(100 * time.Millisecond)
	close(timeTrack)
	sort.Slice(timeStorage, func(i, j int) bool { return timeStorage[i] < timeStorage[j] })
	fmt.Printf("Found %d errors of %d requests for %s. Error ratio: %f. RPS: %f.\n",
		errorNumber, requestNumber, elapsed.String(), float64(errorNumber)/float64(requestNumber), float64(requestNumber)/elapsed.Seconds())
	fmt.Printf("Latency: AVG: %s, MIN: %s, MAX: %s, 50P: %s, 90P: %s, 99P: %s.\n",
		time.Duration(elapsed.Nanoseconds()/requestNumber),
		time.Duration(timeStorage[0]),
		time.Duration(timeStorage[len(timeStorage)-1]),
		time.Duration(timeStorage[len(timeStorage)/2]),
		time.Duration(timeStorage[len(timeStorage)/100*90]),
		time.Duration(timeStorage[len(timeStorage)/100*99]),
	)
	return nil
}

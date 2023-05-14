package concurrency

import (
	"bufio"
	"fmt"
	"go-learinng/utils"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

func ConcurrentWriting() {
	currencies := make(map[string]bool)
	for i := 0; i < 2000; i++ {
		currencies[utils.RandEngString(3)] = true
	}
	startTime := time.Now()
	var counter uint32

	for currency := range currencies {
		err := writeFxRateToFileWithoutBufio(currency, fxRates(currency))
		if err != nil {
			log.Fatal(err)
		}
		err = deleteFxRateFile(currency)
		if err != nil {
			log.Fatal(err)
		}
		counter++
	}
	fmt.Println("Elapsed synchronised without bufio for", counter, "currencies", time.Since(startTime))
	startTime = time.Now()
	counter = 0

	for currency := range currencies {
		err := writeFxRateToFile(currency, fxRates(currency))
		if err != nil {
			log.Fatal(err)
		}
		err = deleteFxRateFile(currency)
		if err != nil {
			log.Fatal(err)
		}
		counter++
	}
	fmt.Println("Elapsed synchronised for", counter, "currencies", time.Since(startTime))

	waitGroup := sync.WaitGroup{}
	startTime = time.Now()
	counter = 0

	for currency := range currencies {
		waitGroup.Add(1)
		currency := currency
		go func() {
			err := writeFxRateToFile(currency, fxRates(currency))
			if err != nil {
				log.Fatal(err)
			}
			err = deleteFxRateFile(currency)
			if err != nil {
				log.Fatal(err)
			}
			atomic.AddUint32(&counter, 1)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	fmt.Println("Elapsed with waitGroup for", counter, "currencies", time.Since(startTime))

	startTime = time.Now()
	currencyChan := make(chan string, len(currencies)/10+5)
	counter = 0
	total := uint32(len(currencies))
	var counterClosed uint32

	for currency := range currencies {
		currency := currency
		go func() {
			err := writeFxRateToFile(currency, fxRates(currency))
			if err != nil {
				log.Fatal(err)
			}
			currencyChan <- currency
			atomic.AddUint32(&counter, 1)
			if atomic.LoadUint32(&counter) == total {
				close(currencyChan)
			}
		}()
	}
	for {
		select {
		case currency, ok := <-currencyChan:
			if ok {
				go func() {
					err := deleteFxRateFile(currency)
					if err != nil {
						log.Fatal(err)
					}
					atomic.AddUint32(&counterClosed, 1)
				}()
			}
		}
		if atomic.LoadUint32(&counterClosed) == total {
			break
		}
	}

	fmt.Println("Elapsed with channels with atomic for", counter, "currencies", time.Since(startTime))

	startTime = time.Now()
	currencyChan = make(chan string, len(currencies)/10+5)
	done := make(chan int)
	counter = 0

	for currency := range currencies {
		currency := currency
		waitGroup.Add(1)
		go func() {
			err := writeFxRateToFile(currency, fxRates(currency))
			if err != nil {
				log.Fatal(err)
			}
			currencyChan <- currency
			waitGroup.Done()
		}()
	}
	go func() {
		waitGroup.Wait()
		close(currencyChan)
		done <- 1
	}()
	removeWithChannels(currencyChan, done)

	fmt.Println("Elapsed with channels with waitGroup for", len(currencies), "currencies", time.Since(startTime))
}

func currencyFileName(currency string) string {
	return filepath.Join(os.TempDir(), currency+"_rate.tmp")
}

func writeFxRateToFile(currency string, rates [365]string) error {
	file, err := os.Create(currencyFileName(currency))
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	for _, rate := range rates {
		_, err = writer.WriteString(rate + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	return err
}

func writeFxRateToFileWithoutBufio(currency string, rates [365]string) error {
	file, err := os.Create(currencyFileName(currency))
	if err != nil {
		return err
	}
	for _, rate := range rates {
		_, err = file.WriteString(rate + "\n")
		if err != nil {
			return err
		}
	}
	return err
}

func deleteFxRateFile(currency string) error {
	return os.Remove(currencyFileName(currency))
}

func fxRates(currency string) [365]string {
	rates := [365]string{}
	basicRate := float64(rand.Int31n(1000))
	for i := 0; i < len(rates); i++ {
		rates[i] = fmt.Sprintf(
			"%s - \"%s\" - %f rate",
			time.Now().Add(time.Hour*24*time.Duration(i-364)).Format("2006-02-01"),
			currency,
			basicRate*(0.9+0.2*rand.Float64()),
		)
	}
	return rates
}

func removeWithChannels(currencies <-chan string, done <-chan int) {
	for {
		select {
		case currency, ok := <-currencies:
			if ok {
				go func() {
					err := deleteFxRateFile(currency)
					if err != nil {
						log.Fatal(err)
					}
				}()
			}
		case <-done:
			return
		}
	}
}

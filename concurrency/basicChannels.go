package concurrency

import (
	"fmt"
	"time"
)

func BasicChannels() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		for i := 0; i < 18; i++ {
			time.Sleep(200 * time.Millisecond)
			channel1 <- fmt.Sprintf("Channel 1 incremented to %d after 200 mills", i)
		}
		close(channel1)
	}()
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			channel1 <- fmt.Sprintf("Channel 2 incremented to %d after 1 second", i)
		}
		close(channel2)
	}()
	for {
		var msg1, msg2 string
		var open1, open2 bool
		select {
		case msg1, open1 = <-channel1:
			fmt.Println(msg1)
		case msg2, open2 = <-channel2:
			fmt.Println(msg2)
		}
		if !(open1 || open2) {
			break
		}
	}
}

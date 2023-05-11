package concurrency

import (
	"context"
	"fmt"
	"time"
)

func Contexts() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "id", 8534)
	handle(ctx)
}

func handle(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()
	process(ctx)
	fmt.Printf("Handling of %d request has finnished\n", ctx.Value("id"))
}

func process(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("Done!")
			return
		case <-ctx.Done():
			fmt.Println("Timeout!")
			return
		}
	}
}

package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerCh := make(chan string)
	doneCh := make(chan struct{})

	go ping(ctx, pingerCh)
	go pong(ctx, pingerCh)

	go func() {
		timeout := time.After(5 * time.Second)

		for {
			select {
			case <-timeout:
				fmt.Println("timenout")
				close(pingerCh)
				doneCh <- struct{}{}
				return

			case msg := <-pingerCh:
				fmt.Println("ping received:", msg)
			}
		}
	}()

	<-doneCh
	fmt.Println("done")
}

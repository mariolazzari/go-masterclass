package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	// done := make(chan bool)

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Receving:", r)
			} else {
				fmt.Println("Close channel")
				//done <- true
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending:", i)
	}
	close(jobs)

	wg.Wait()
	// <-done
}

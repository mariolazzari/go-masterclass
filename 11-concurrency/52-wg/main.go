package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(msg string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	fmt.Println("Main start")

	go sayHello("Ciao", time.Second, &wg)
	go sayHello("Ciao ciao", time.Second, &wg)
	go sayHello("Ciao 2s", 2*time.Second, &wg)
	go sayHello("Ciao 3s", 3*time.Second, &wg)

	totJobs := 5
	for i := range totJobs {
		wg.Add(1)
		go sayHello(fmt.Sprintf("Job %d", i), time.Second, &wg)
	}

	fmt.Println("Main end")

	wg.Wait()
}

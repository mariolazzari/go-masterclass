package main

import (
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", msg)
}

func main() {
	fmt.Println("Main start")

	go sayHello("Ciao 1s", time.Second)
	go sayHello("Ciao 2s", 2*time.Second)
	go sayHello("Ciao 3s", 3*time.Second)

	fmt.Println("Main end")
	time.Sleep(2 * time.Second)
}

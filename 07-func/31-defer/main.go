package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Start simpleDefer")
	defer fmt.Println("1st Defer simpleDefer")
	defer fmt.Println("2nd Defer simpleDefer")
	fmt.Println("Middle simpleDefer")
}

func main() {
	defer func() {
		fmt.Println("Before main returns")
	}()
	simpleDefer()

	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("Main last")
}

package main

import "fmt"

func mayPanic(isPanic bool) {
	if isPanic {
		panic("panic!")
	}
	fmt.Println("Done")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic")
		}
	}()

	mayPanic(true)
}

func main() {
	mayPanic(false)
	recoverable()
}

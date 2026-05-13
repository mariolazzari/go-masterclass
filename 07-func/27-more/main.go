package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	f := factorial(5)
	fmt.Println(f)

	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Mario")
}

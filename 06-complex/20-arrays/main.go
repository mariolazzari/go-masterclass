package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("Numbers: %+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("Numbers: %+v\n", numbers)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Primes: %+v\n", primes)

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 3
	fmt.Printf("Matrix: %+v\n", matrix)
}

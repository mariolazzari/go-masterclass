package main

import "fmt"

type Number interface {
	int | float32 | float64 | string
}

func sum[T Number](nums ...T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	names := []string{"Mario", "Maria", "Mariarosa"}
	fmt.Println("Total names:", sum(names...))

	grades := []int{10, 20}
	fmt.Println("Total grades:", sum(grades...))

	floats := []float64{1.23, 2.34, 3.45}
	fmt.Println("Total floats:", sum(floats...))
}

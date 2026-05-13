package main

import "fmt"

func greet(name string) {
	fmt.Printf("Ciao %s\n", name)
}

func add(a, b int) int {
	return a + b
}

func area(width, height float64) float64 {
	if width < 0 || height < 0 {
		return 0
	}

	return width * height
}

func main() {
	greet("Mario")

	a := 1
	b := 2
	c := add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, c)

	fmt.Println("area = ", area(4, 5))

}

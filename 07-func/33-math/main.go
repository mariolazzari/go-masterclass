package main

import (
	"fmt"
	"strings"
)

const (
	division  = "Division"
	divByZero = "Division by 0"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

func (e MathError) Error() string {
	var inputs []string

	if e.Operation == "Division" {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s", e.Operation, strings.Join(inputs, ","), e.Message)
}

func sum(nums ...int) int {
	defer fmt.Println("Sum finished")

	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func safeDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divByZero,
		}
	}

	return a / b, nil
}

func main() {
	fmt.Printf("sum: %d\n", sum(1, 2, 3, 4))
	res, err := safeDiv(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

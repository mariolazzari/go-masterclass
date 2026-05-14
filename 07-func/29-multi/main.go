package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}

	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}

func main() {
	res, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		println(res)
	}

	res, err = divide(4, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		println(res)
	}

	firstName, lastName := splitName("Mario Lazzari")
	fmt.Printf("First name: %s\nLast name: %s\n", firstName, lastName)
}

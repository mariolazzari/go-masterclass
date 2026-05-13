package main

import "fmt"

func modifyValue(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}

	*val *= 10
	fmt.Println("modifyValue:", *val)
}

func main() {
	age := 51
	ptrAge := &age

	fmt.Println("age    :", age)
	fmt.Println("ptrAge :", ptrAge)
	fmt.Println("*ptrAge:", *ptrAge)

	modifyValue(&age)
	fmt.Println("main", age)
	modifyValue(nil)
}

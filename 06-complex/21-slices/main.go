package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Mario", "Mariarosa", "Maria"}
	fmt.Printf("Names: %+v\n", names)

	items := make([]int, 3, 5)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	fmt.Println("items[1:5]", items[1:5])
	fmt.Println("items[:5]", items[:5])
	fmt.Println("items[5:]", items[5:])
	fmt.Println("items[5:]", items[:])

	fmt.Println("Advance slice")
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("slice: %+v, Len: %d, Cap: %d\n", slice, len(slice), cap(slice))

	s1 := slice[2:5]
	fmt.Printf("s1: %+v, Len: %d, Cap: %d\n", s1, len(s1), cap(s1))

	s2 := slice[:4]
	fmt.Printf("s2: %+v, Len: %d, Cap: %d\n", s2, len(s2), cap(s2))

	s3 := slice[6:]
	fmt.Printf("s3: %+v, Len: %d, Cap: %d\n", s3, len(s3), cap(s3))

	s4 := slice[:]
	fmt.Printf("s4: %+v, Len: %d, Cap: %d\n", s4, len(s4), cap(s4))

	ok := slices.Contains(slice, 4)
	if ok {
		fmt.Println("4 is present")
	}

	s4 = append(s4, 1000)
	fmt.Printf("s4: %+v, Len: %d, Cap: %d\n", s4, len(s4), cap(s4))

}

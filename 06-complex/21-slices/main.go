package main

import "fmt"

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

}

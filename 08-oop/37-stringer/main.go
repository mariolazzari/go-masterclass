package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func (e Employee) String() string {
	return fmt.Sprintf("%s (%d)", e.Name, e.ID)
}

type ID int

func (id ID) String() string {
	return fmt.Sprintf("My id is: %d", id)
}

func main() {
	mario := Employee{
		ID:   1,
		Name: "Mario",
	}
	fmt.Println(mario)

	myID := ID(1)
	fmt.Println(myID)
}

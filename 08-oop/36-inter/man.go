package main

import "fmt"

type Person interface {
	GetName() string
}

type Employee struct {
	ID   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

type Business struct {
	ID   int
	Name string
}

func (b Business) GetName() string {
	return b.Name
}

func displayPerson(e Person) {
	fmt.Println(e.GetName())

}

func main() {
	mario := Employee{
		ID:   1,
		Name: "Mario",
	}
	displayPerson(mario)

	maria := Business{
		ID:   1,
		Name: "Maria",
	}
	displayPerson(maria)
}

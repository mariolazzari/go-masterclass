package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, position string, salary int) *Employee {
	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Salary:    salary,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
}

func main() {

	mario := Employee{
		ID:        1,
		FirstName: "Mario",
		LastName:  "Lazzari",
		Position:  "Full stack developer",
		Salary:    1000000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Printf("Mario: %+v\n", mario)
	fmt.Println(mario.FirstName)

	maria := NewEmployee(2, "Maria", "Lazzari", "Teacher", 2000000)
	fmt.Printf("Maria: %+v\n", maria)
	fmt.Println(maria.FirstName)

	mario.Salary *= 2
	fmt.Printf("Mario salary: %d", mario.Salary)

}

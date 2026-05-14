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

func (e Employee) FullName() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}

func (e *Employee) Deactivate() {
	e.IsActive = false
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

	fmt.Println(mario.FullName())

	mario.Deactivate()
	fmt.Println(mario.IsActive)

}

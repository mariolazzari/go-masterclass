# Go (Golang) Masterclass: Learn Like a Google Engineer

[Udemy](https://www.udemy.com/course/learn-golang-like-google-engineers-do)

## Setup

### First program

[Github](https://github.com/joefazee/learning-go)
[Postman](https://www.postman.com/dark-shuttle-458212/learning-go-shop/overview?sideView=agentMode)

```go
package main

import "fmt"

func main() {
	fmt.Println("Ciao")
}
```

### Go toolchain

```sh
go run main.go
go build main.go
go fmt main.go
go test .
```

## Setup

### Links

[Go migrate](https://github.com/golang-migrate/migrate)
[Linter](https://golangci-lint.run/)

## Core

### Variables

```go
package main

import "fmt"

func main() {

package main

import "fmt"

func main() {

	var greeting string // zero-value is an empty string ""
	greeting = "Hello, world!"

	fmt.Println(greeting)

	var count int // zero-value 0
	count = 10
	fmt.Println(count)

	var isRunning bool // zero-value false
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Println(firstName, lastName)

	// Short declaration
	email := "test@test.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	var year = 2025
	fmt.Println(year)

}
```

### Constants

```go
package main

import "fmt"

const (
	Host = "127.0.0.1"
	Port = ":8080"
	User = "root"
)

var (
	isRunning bool
)

func main() {

	AppName := "Go"
	fmt.Println(AppName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 5.2
	fmt.Println(rate)

}
```

### Enums

```go
package main

import "fmt"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
```

### Project: Custom logger

```go
package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}

	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}
```

## Control flows

### For loop

```go
package main

import "fmt"

func main() {

	// for -- only way to loop

	// C-style loop
	for i := 1; i <= 10; i++ {
		//fmt.Println(i)
	}

	// while-style
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}
	fmt.Println("------------ infinite loop ------------")
	counter := 0
	for {
		fmt.Println("counter:", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	fmt.Println("------------ skipping---------")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("------------ array---------")
	items := [3]string{"Go", "TypeScript", "Rust"}
	for index, _ := range items {
		fmt.Println(items[index])
	}

}
```

### If

```go
package main

import "fmt"

func main() {

	tmp := 25
	if tmp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("greater is less than 30")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C, B")
	} else {
		fmt.Println("Failed")
	}

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["john"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("access not granted")
	}

}
```

### Switch

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Monday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i any) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.23)

}
```

### Project: sales order

```go
package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]
	if !found {
		if before, ok := strings.CutSuffix(itemCode, "_SALE"); ok {
			originalItemCode := before
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original: $%.2f, Sale Price: $%.2f)\n",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}

		fmt.Printf(" - Item: %s (Product not found)\n", itemCode)
		return 0.0, false
	}

	return basePrice, true

}

func main() {
	fmt.Println("-------------- Simple Sales Order Processor------------")

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	var subtotal float64
	fmt.Println("-------Processing Order Items:")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Subtotal Price: %.2f\n", subtotal)
}
```

## Complex data types

### Array

```go
package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("Numbers: %+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("Numbers: %+v\n", numbers)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Primes: %+v\n", primes)

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 3
	fmt.Printf("Matrix: %+v\n", matrix)
}
```

### Slices

```go
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
```

### Maps

```go
package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Mario":     30,
		"Mariarosa": 30,
		"Gino":      18,
	}
	fmt.Printf("Grades: %+v\n", studentGrades)

	studentGrades["Gino"] = 20
	fmt.Printf("Grades: %+v\n", studentGrades)

	mario, ok := studentGrades["Mario"]
	if ok {
		fmt.Printf("Mario: %d\n", mario)
	}

	_, ok = studentGrades["Pino"]
	if ok {
		fmt.Println("Pino was here")
	}

	delete(studentGrades, "Gino")
	fmt.Printf("Grades: %+v\n", studentGrades)

	configs := make(map[string]int)
	fmt.Printf("%+v, %T\n", configs, configs)

	if configs == nil {
		fmt.Println("configs is nil")
	}

	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
```

### Pointers

```go
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
```

### Project: CMS

```go
package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact %s exists\n", name)
		return
	}

	contactList = append(contactList, Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	})
	contactIndexByName[name] = nextID

	fmt.Printf("Contact %s added\n", name)
	nextID++
}

func findContact(name string) *Contact {
	if idx, ok := contactIndexByName[name]; ok {
		return &contactList[idx]
	}
	return nil
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No contact found")
	}

	for _, c := range contactList {
		fmt.Println(c)
	}
}

func main() {
	listContacts()
	addContact("Mario", "mario@lazzari.it", "123-1234")
	addContact("Maria", "maria@lazzari.it", "123-4321")
	listContacts()

	mario := findContact("Mario")
	if mario == nil {
		fmt.Println("Mario not found")
	} else {
		fmt.Println("Found:", *mario)
	}
}
```

## Functions & errors

### Functions

```go
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
```

### More on functions

```go
package main

import "fmt"

func factorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	f := factorial(5)
	fmt.Println(f)

	next := intSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Mario")
}
```

### Variadic functions

```go
package main

import "fmt"

func add(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	var sum int = 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

func configs(nums ...int) {
	if len(nums) > 0 {
		fmt.Println(nums[0])
	} else {
		fmt.Println("Default")
	}

}

func main() {
	fmt.Println(add())
	fmt.Println(add(1))
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))

	configs()
	configs(1, 2, 3)
}
```

### Multi values return

```go
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
```

### Custon error

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivByZero = errors.New("division by 0")
var ErrTooLarge = errors.New("a too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func (o OpError) Error() string {
	return o.Message
}

func NewOpError(op string, code int, msg string, time time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: msg,
		Time:    time,
	}
}

func DoError() error {
	return NewOpError("not found", 404, "Not found", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}

	if a > 1000 {
		return 0, ErrTooLarge
	}

	return a / b, nil
}

func main() {
	res, err := divide(4, 0)
	if err != nil {
		if errors.Is(err, ErrDivByZero) {
			fmt.Println("b cannot be 0")
		}
	} else {
		println(res)
	}

	res, err = divide(4000, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		println(res)
	}

	err = DoError()
	fmt.Printf("%s", err)
}
```

### Defer

```go
package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Start simpleDefer")
	defer fmt.Println("1st Defer simpleDefer")
	defer fmt.Println("2nd Defer simpleDefer")
	fmt.Println("Middle simpleDefer")
}

func main() {
	defer func() {
		fmt.Println("Before main returns")
	}()
	simpleDefer()

	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("Main last")
}
```

### Panic

```go
package main

import "fmt"

func mayPanic(isPanic bool) {
	if isPanic {
		panic("panic!")
	}
	fmt.Println("Done")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic")
		}
	}()

	mayPanic(true)
}

func main() {
	mayPanic(false)
	recoverable()
}
```

### Project math

```go
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
```

## OOP in Go

### Custom types

```go
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
```

### Methods

```go
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
```

### Interfaces

```go
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
```

### Stringer interface

[Stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)

```go
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
```

### Generics

```go
package main

import "fmt"

type Number interface {
	int | float32 | float64 | string
}

func sum[T Number](nums ...T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	names := []string{"Mario", "Maria", "Mariarosa"}
	fmt.Println("Total names:", sum(names...))

	grades := []int{10, 20}
	fmt.Println("Total grades:", sum(grades...))

	floats := []float64{1.23, 2.34, 3.45}
	fmt.Println("Total floats:", sum(floats...))
}
```

### Project: payroll processor

```go
package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64 // Calculates monthly pay
	// We'll also rely on fmt.Stringer, so our types should implement String()
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64 // Hours worked in the month
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64 // Monthly base
	CommissionRate float64 // e.g., 0.05 for 5%
	SalesAmount    float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%.2f)",
		ce.Name, ce.BaseSalary, ce.CommissionRate*100, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("  - Processing: %s\n", employee) // Relies on String() method
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0
	for _, emp := range employees {
		PrintEmployeeSummary(emp) // Generic function call
		pay := emp.CalculatePay()
		fmt.Printf("    Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("--------------------------")
}

func main() {
	fmt.Println("Welcome to the Payroll Processor!")
	salEmp := SalariedEmployee{Name: "Alice Wonderland", AnnualSalary: 72000.00}
	hrEmp := HourlyEmployee{Name: "Bob The Builder", HourlyRate: 25.00, HoursWorked: 160.0}
	comEmp := CommissionEmployee{Name: "Charlie Chaplin", BaseSalary: 2000.00, CommissionRate: 0.10, SalesAmount: 15000.00}

	payrollList := []Payable{
		salEmp,
		hrEmp,
		comEmp,
		HourlyEmployee{Name: "Diana Prince", HourlyRate: 30.00, HoursWorked: 150.0},
	}

	ProcessPayroll(payrollList)
}
```

## Composition design pattern

### Composition

```go
package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Mail            string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Println("Customer ID     :", c.CustomerID)
	fmt.Println("Customer Name   :", c.Name)
	fmt.Println("Customer eMail  :", c.Mail)
	fmt.Println("Billing Address :", c.CustomerID)
	fmt.Println("Shipping Address:", c.ShippingAddress.FullAddress())
	fmt.Println("Billing  Address:", c.BillingAddress.FullAddress())
}

func main() {
	cust := Customer{
		CustomerID: 1,
		Name:       "Mario",
		Mail:       "mario.lazzari@gmail.com",
		BillingAddress: Address{
			Street:  "stree 1",
			City:    "City 1",
			State:   "State 1",
			ZipCode: "12345",
		},
		ShippingAddress: Address{
			Street:  "stree 2",
			City:    "City 2",
			State:   "State 2",
			ZipCode: "54321",
		},
	}

	cust.PrintDetails()
}
```

### Embedding over inheritance

```go
package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.ZipCode)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("%s, %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Println("Company Name:", c.Name)
	fmt.Println("Address     :", c.FullAddress())
	fmt.Println("Company Street (promoted):", c.Street)

}

type CompanyWithOwnMail struct {
	Company
	Email string // shadowing
}

func main() {
	comp := Company{
		Name: "Company 1",
		Address: Address{
			City:    "City 1",
			Street:  "Street 1",
			State:   "State 1",
			ZipCode: "12345",
		},
		BusinessType: "Type 1",
	}

	comp.GetProfile()

	comp2 := CompanyWithOwnMail{
		Company: Company{
			Name: "Company 12",
			ContactInfo: ContactInfo{
				Email: "original@mail.com",
			},
			Address: Address{
				City:    "City 1",
				Street:  "Street 1",
				State:   "State 1",
				ZipCode: "12345",
			},
			BusinessType: "Type 1",
		},
		Email: "shadowing@mail.com",
	}

	comp2.GetProfile()
	fmt.Println("Shadowing Email:", comp2.Email)

}
```

### Project: bank account

```go
package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited $%.2f to %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return fmt.Errorf("insufficient funds in %s. Balance: $%.2f, Tried to withdraw: $%.2f",
			acc.AccountNumber, acc.Balance, amount)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s, Balance: $%.2f",
		acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account              // Embed Account struct (anonymous field)
	InterestRate float64 // e.g., 0.02 for 2%
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate // Accesses promoted Balance field
	fmt.Printf("Adding interest $%.2f to savings account %s. ", interest, sa.AccountNumber)
	err := sa.Deposit(interest) // Uses promoted Deposit method
	if err != nil {
		fmt.Printf("AddInterst: Error depositing $%.2f to savings account. %v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account        // Embed Account struct
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	// Allow withdrawal up to Balance + OverdraftLimit
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %s. Available including overdraft: $%.2f",
			amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount // Balance can go negative
	fmt.Printf("Withdrew $%.2f from overdraft account %s. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {

	fmt.Println("--- Bank Account System ---")

	savAcc := SavingsAccount{
		Account: Account{ // Initialize the embedded Account
			AccountNumber: "SAV001",
			Balance:       1000.00,
			OwnerName:     "Alice Saver",
		},
		InterestRate: 0.02, // 2%
	}
	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200.00)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %v\n", 200.00, err)
	}
	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := OverdraftAccount{
		Account: Account{
			AccountNumber: "OVD002",
			Balance:       100.00,
			OwnerName:     "Bob Spender",
		},
		OverdraftLimit: 200.00,
	}

	fmt.Println("\n--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())
}
```

## Strings

### Documentation

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"
	s2 := strings.Clone(s1)
	fmt.Printf("s1=%s, s2=%s\n", s1, s2)

	b := strings.Builder{}
	b.Write([]byte("byte example"))
	fmt.Println(b.String())
	b.WriteString(", string example\n")
	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.Title(s1))

	s3 := "   " + s1 + "     "
	fmt.Println(len(s3))
	s3 = strings.TrimSpace(s3)
	fmt.Println(len(s3))

	s3 = "mario.lazzari@gmail.com"
	fmt.Println(strings.HasPrefix(s3, "mario"))
	fmt.Println(strings.HasSuffix(s3, "gmail.com"))
	fmt.Println(strings.Replace(s3, "mario", "maria", 1))

	parts := strings.Split(s3, "@")
	fmt.Printf("%+v\n", parts)
	parts = strings.Fields("Mario Lazzari")
	fmt.Printf("%+v\n", parts)

	fmt.Println(strings.Join(parts, ","))
}
```

### Formatting strings

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

type Config struct {
	Key   string
	Value any
	IsSet bool
}

func (c Config) String() string {
	return fmt.Sprintf("Key %s, Value %s, IsSet %t\n", c.Key, c.Value, c.IsSet)
}

func main() {
	cfg := Config{
		Key:   "key",
		Value: 10,
		IsSet: true,
	}

	// default formatting
	fmt.Printf("%v", cfg)
	// with field names
	fmt.Printf("%+v", cfg)
	// with value
	fmt.Printf("%#v\n", cfg)
	// with type
	fmt.Printf("%T\n", cfg)

	// errors
	port := 3000
	err := errors.New("port already in use")
	err = fmt.Errorf("error on port %d: %w", port, err)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Unicode

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	mario := "mario"
	marioJap := "マリオ"

	fmt.Println(len(mario))
	fmt.Println(len(marioJap))

	fmt.Printf("%c\n", mario[0])
	fmt.Printf("%c\n", mario[0])

	for _, c := range mario {
		fmt.Println(c)
	}
	for _, c := range mario {
		fmt.Println(string(c))
	}

	for _, c := range marioJap {
		fmt.Println(c)
	}
	for _, c := range marioJap {
		fmt.Println(string(c))
	}

	// rune
	fmt.Println("rune")
	data := []rune{'マ', 'リ', 'オ'}
	for _, d := range data {
		fmt.Println(string(d), unicode.IsLower(d), unicode.IsLetter(d))
	}

}
```

### Regex

```go
package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello, World! Welcome to Go!"

	regGo, err := regexp.Compile(`Go`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text '%s' matches 'Go': %t\n", text1, regGo.MatchString(text1))

	text2 := "Product codes: P123, X123, P789"

	regProd := regexp.MustCompile(`P\d+`)

	firstProd := regProd.FindString(text2)
	fmt.Println(string(firstProd))

	allProds := regProd.FindAllString(text2, -1)
	fmt.Println(allProds)
}
```

### Templates

```go
package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string // demo a loop
	UnreadCount   int
}

func main() {

	fmt.Println("--- Text template example ---")

	emailTemplate := `
Subject: {{ .Subject }}

{{.Body}}

{{if .Items}}
   Related Items:
{{range .Items}}
	- {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnreadCount}} unreads.
{{else}}
You have no messages
{{end}}


- Thanks
{{.SenderName}}
`
	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecipientName: "Alice",
		SenderName:    "Bob's Auto-Responder",
		Subject:       "Your Weekly Update",
		Body:          "Here is the update you requested. We hope you find it useful.",
		Items:         []string{"Report A", "Document B", "Summary C"},
		UnreadCount:   0,
	}

	var output strings.Builder

	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(strings.ToUpper(output.String()))
}
```

### Project: config parser

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseConfig(content string) (map[string]string, error) {
	config := make(map[string]string)

	re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)

	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		matches := re.FindStringSubmatch(trimmedLine)
		if matches == nil {
			fmt.Printf("Line %d: '%s' - Is Inavlid\n", lineNo, line)
			continue
		}
		key := matches[1]
		var value string

		if matches[2] != "" {
			value = matches[2]
		} else if matches[3] != "" {
			value = matches[3]
		} else {
			value = matches[4]
		}

		config[key] = value
	}

	return config, nil
}

func main() {

	envFileContent := `
# Application Configuration
APP_NAME="My Cool App"
APP_VERSION="1.0.2-beta" # Version with quotes
PORT=8080
DEBUG_MODE="true"
# Database Settings
DB_HOST=localhost
DB_USER = admin
DB_PASSWORD = "p@s$w Ord With Sp@ces!" # Quoted password
API_ENDPOINT = https://api.example.com/v1

# An empty value
EMPTY_KEY=
ANOTHER_KEY_NO_VALUE =`

	config, err := parseConfig(envFileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for k, v := range config {
		fmt.Printf("%s=%q\n", k, v)
	}

}
```

## Modules

### Go module

```sh
go mod init github.com/mariolazzari/my-module
```

## Concurrency

### Starting go routine

```go
package main

import (
	"fmt"
	"time"
)

func sayHello(msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", msg)
}

func main() {
	fmt.Println("Main start")

	go sayHello("Ciao 1s", time.Second)
	go sayHello("Ciao 2s", 2*time.Second)
	go sayHello("Ciao 3s", 3*time.Second)

	fmt.Println("Main end")
	time.Sleep(2 * time.Second)
}
```

### Wait groups

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(msg string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	fmt.Println("Main start")

	go sayHello("Ciao", time.Second, &wg)
	go sayHello("Ciao ciao", time.Second, &wg)
	go sayHello("Ciao 2s", 2*time.Second, &wg)
	go sayHello("Ciao 3s", 3*time.Second, &wg)

	totJobs := 5
	for i := range totJobs {
		wg.Add(1)
		go sayHello(fmt.Sprintf("Job %d", i), time.Second, &wg)
	}

	fmt.Println("Main end")

	wg.Wait()
}
```

### Channels

```go
package main

import "fmt"

type user struct {
	name string
}

func main() {
	messages := make(chan string)
	users := make(chan user)

	go func() {
		fmt.Println("Sending message to messages channel")
		messages <- "Ciao"
	}()

	go func() {
		fmt.Println("Sending message to users channel")
		users <- user{
			name: "Mario",
		}
	}()

	msg := <-messages
	fmt.Println(msg)

	user := <-users
	fmt.Println(user)

}
```

### Buffered channels

```go
package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	fmt.Println("Send message to buffered channel")
	messages <- "Msg1"
	messages <- "Msg2"
	messages <- "Msg3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
```

### Closing channel

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 5)
	// done := make(chan bool)

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Receving:", r)
			} else {
				fmt.Println("Close channel")
				//done <- true
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending:", i)
	}
	close(jobs)

	wg.Wait()
	// <-done
}
```

### Project: ping-ponger

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerCh := make(chan string)
	doneCh := make(chan struct{})

	go ping(ctx, pingerCh)
	go pong(ctx, pingerCh)

	go func() {
		timeout := time.After(5 * time.Second)

		for {
			select {
			case <-timeout:
				fmt.Println("timenout")
				close(pingerCh)
				doneCh <- struct{}{}
				return

			case msg := <-pingerCh:
				fmt.Println("ping received:", msg)
			}
		}
	}()

	<-doneCh
	fmt.Println("done")
}
```

### Project: concurrent downloader

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pingerCh := make(chan string)
	doneCh := make(chan struct{})

	go ping(ctx, pingerCh)
	go pong(ctx, pingerCh)

	go func() {
		timeout := time.After(5 * time.Second)

		for {
			select {
			case <-timeout:
				fmt.Println("timenout")
				close(pingerCh)
				doneCh <- struct{}{}
				return

			case msg := <-pingerCh:
				fmt.Println("ping received:", msg)
			}
		}
	}()

	<-doneCh
	fmt.Println("done")
}
```

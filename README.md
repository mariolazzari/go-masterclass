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

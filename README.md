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

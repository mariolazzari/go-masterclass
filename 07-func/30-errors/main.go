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

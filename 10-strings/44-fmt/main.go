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

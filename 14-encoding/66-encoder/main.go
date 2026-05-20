package main

import (
	"encoding/json"
	"log"
	"os"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone"`
	Password string `json:"-" xml:"-"`
	IsActvie bool   `json:"is_active" xml:"is_active"`
}

func main() {
	mario := user{
		Name:     "Mario",
		Age:      51,
		Phone:    "123-1234",
		IsActvie: true,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&mario); err != nil {
		log.Fatal(err)
	}
}

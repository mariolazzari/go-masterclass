package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone"`
	IsActvie bool   `json:"is_active" xml:"is_active"`
}

func main() {
	mario := user{
		Name:     "Mario",
		Age:      51,
		Phone:    "123-1234",
		IsActvie: true,
	}

	// JSON
	byteSlice, err := json.Marshal(mario)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteSlice))

	// JSON with identation
	byteSlice, err = json.MarshalIndent(mario, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteSlice))

	// XML
	byteSlice, err = xml.Marshal(mario)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteSlice))

}

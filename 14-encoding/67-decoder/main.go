package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone"`
	Password string `json:"-" xml:"-"`
	IsActvie bool   `json:"is_active" xml:"is_active"`
}

var payload = `{
  "name": "Mario",
  "age": 51,
  "phone": "123-1234",
  "password": "s3cret",
  "is_active": true,
  "profile": {
  	"url": "mariolazzari.it"
  }
}`

func main() {
	var u user
	enc := json.NewDecoder(strings.NewReader(payload))
	if err := enc.Decode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

}

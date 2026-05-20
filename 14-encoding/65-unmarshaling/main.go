package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"name" xml:"name"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phone" xml:"phone"`
	Password string  `json:"-" xml:"-"`
	IsActvie bool    `json:"is_active" xml:"is_active"`
	Profile  profile `json:"profile" xml:"profile"`
}

type profile struct {
	URL string `json:"url" xml:"url"`
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
	var mario user
	err := json.Unmarshal([]byte(payload), &mario)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mario: %+v\n", mario)

	m, err := json.MarshalIndent(mario, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("m: %s\n", m)

}

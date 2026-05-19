package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
)

//go:embed hello.txt
var data string

//go:embed public
var public embed.FS

func main() {
	fmt.Println(data)

	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}

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

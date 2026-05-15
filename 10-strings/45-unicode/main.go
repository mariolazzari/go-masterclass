package main

import (
	"fmt"
	"unicode"
)

func main() {
	mario := "mario"
	marioJap := "マリオ"

	fmt.Println(len(mario))
	fmt.Println(len(marioJap))

	fmt.Printf("%c\n", mario[0])
	fmt.Printf("%c\n", mario[0])

	for _, c := range mario {
		fmt.Println(c)
	}
	for _, c := range mario {
		fmt.Println(string(c))
	}

	for _, c := range marioJap {
		fmt.Println(c)
	}
	for _, c := range marioJap {
		fmt.Println(string(c))
	}

	// rune
	fmt.Println("rune")
	data := []rune{'マ', 'リ', 'オ'}
	for _, d := range data {
		fmt.Println(string(d), unicode.IsLower(d), unicode.IsLetter(d))
	}

}

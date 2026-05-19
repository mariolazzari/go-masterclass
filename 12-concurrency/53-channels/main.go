package main

import "fmt"

type user struct {
	name string
}

func main() {
	messages := make(chan string)
	users := make(chan user)

	go func() {
		fmt.Println("Sending message to messages channel")
		messages <- "Ciao"
	}()

	go func() {
		fmt.Println("Sending message to users channel")
		users <- user{
			name: "Mario",
		}
	}()

	msg := <-messages
	fmt.Println(msg)

	user := <-users
	fmt.Println(user)

}

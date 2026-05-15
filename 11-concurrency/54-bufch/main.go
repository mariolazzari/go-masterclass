package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	fmt.Println("Send message to buffered channel")
	messages <- "Msg1"
	messages <- "Msg2"
	messages <- "Msg3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

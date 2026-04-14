package main

import "fmt"

func conditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("trying to send a message of length:", messageLen, "and and max length of:", maxMessageLen)

	// Conditions not within parentheses
	if messageLen < maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}

func main() {
	conditionals()
}

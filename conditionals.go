package main

import (
	"fmt"
	"unicode/utf8"
)

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

func initialStatement() {
	email := "hi@me.com"
	if length := utf8.RuneCountInString(email); length < 10 {
		fmt.Printf("Email must be at least 10 characters , it is %d\n", length)
	}
}

func main() {
	// conditionals()
	initialStatement()
}

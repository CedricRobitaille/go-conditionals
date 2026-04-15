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

	// We can use an initial statement in the `if` statement.
	// This statement can be used in the scope of the `if` body
	// This allows for global variables to exist without interconnection
	if length := utf8.RuneCountInString(email); length < 10 {
		fmt.Printf("Email must be at least 10 characters , it is %d\n", length)
	}
}

// Billing cost function for the switchStatement()
func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}

func switchStatement() {
	plan := "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}

func main() {
	// conditionals()
	// initialStatement()
	switchStatement()
}

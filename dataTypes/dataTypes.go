package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets uint = 100
	var remainingTickets uint = 100

	// when we use variable then we can use it in the expression
	// consferernceName := "Golang Conference"
	// but when we use const it can't use in the expression

	// formatting out-put
	fmt.Printf("%s %d %d\n", "Total tickets available:", conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v let's enjoy go\n", conferenceName)
	fmt.Printf("Welcome to our %s let's enjoy go\n", conferenceName)
	fmt.Printf("Our total ticket %v and remaining ticket %v", conferenceTickets, remainingTickets)

	// formatting input
	fmt.Printf("Enter your name \n: ")
	var userName string
	fmt.Scanf("%s", &userName)
	fmt.Printf("User name %s\n", userName)

	// formatting input
	var userFirstName string
	var userLastName string
	var email string
	var userTickets uint
	fmt.Printf("Enter your First name \n: ")
	fmt.Scanf("%s", &userFirstName)
	fmt.Printf("Enter your Last name \n: ")
	fmt.Scanf("%s", &userLastName)
	fmt.Printf("Enter your email \n: ")
	fmt.Scanf("%s", &email)
	fmt.Printf("Enter your tickets \n: ")
	fmt.Scanf("%d", &userTickets)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email by %v\n", userFirstName, userLastName, userTickets, email)

}

// go is statically type language
// go is compiled statically
// you need to tell go compiler what type of variable you are using
// go compiler will check the type of variable and give you error if you are using wrong type when compile time not run time
// unsigned int is uint and signed int is int and float is float32 and double is float64 and complex is complex64 and complex128
// unit is used for positive value eg. o to ....

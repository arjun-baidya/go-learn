package main

import "fmt"

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets uint = 100
	var remainingTickets uint = 100

	// when we use variable then we can use it in the expression
	// consferernceName := "Golang Conference"
	// but when we use const it can't use in the expression

	// formatting out-put
	fmt.Printf("Welcome to our %v let's enjoy go\n", conferenceName)
	fmt.Printf("Our total ticket %v and remaining ticket %v \n", conferenceTickets, remainingTickets)

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email by %v\n", userFirstName, userLastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)
}

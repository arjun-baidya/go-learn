package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets uint = 100
	var remainingTickets uint = 100
	// slice
	booking := []string{}

	// function call
	greetUasers(conferenceName, conferenceTickets, remainingTickets)

	for {
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

		// validate user input
		isvalidName := len(userFirstName) >= 2 && len(userLastName) >= 2
		isvalidEmail := strings.Contains(email, "@")
		isvalidTicketNumer := userTickets > 0 && userTickets <= remainingTickets
		// here validate user can't take more than remaining tickets

		if isvalidName && isvalidEmail && isvalidTicketNumer {

			remainingTickets = remainingTickets - userTickets
			// slice
			booking = append(booking, userFirstName+" "+userLastName)
			fmt.Printf("The whole slice: %v\n", booking)
			fmt.Printf("The first value of slice %v \n", booking[0])
			fmt.Printf("Slice type %T\n", booking)
			fmt.Printf("slice length %v\n", len(booking))

			fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email by %v\n", userFirstName, userLastName, userTickets, email)
			fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)

			// for loop and nested for loop
			firstNames := []string{}
			// blank identifier is used to ignore the value of the variable _
			for _, booking := range booking {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
			fmt.Printf("First name of bookings %v\n", firstNames)
			// if else statement
			if remainingTickets == 0 {
				fmt.Println("our ticket is booked , Please wait next time")
				break
			}
		} else {
			if !isvalidName {
				fmt.Println("Please enter valid name")
			}
			if !isvalidEmail {
				fmt.Println("Please enter valid email")
			}
			if !isvalidTicketNumer {
				fmt.Println("Please enter valid ticket number")
			}
		}
	}
}

func greetUasers(confName string, confTicket uint, remainingTicket uint) {
	fmt.Printf("Hello welcome to our %v\n", confName)
	fmt.Printf("Our total ticket %v and remaining ticket %v \n", confTicket, remainingTicket)
}

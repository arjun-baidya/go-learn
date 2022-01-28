package main

import (
	"fmt"
	"strconv"
	"strings"
)

var conferenceName string = "Golang Conference"

const conferenceTickets uint = 100

var remainingTickets uint = 100
var booking = make([]map[string]string, 0)

func main() {

	// function call
	greetUasers()

	for {
		userFirstName, userLastName, email, userTickets := getUserInput()
		// here used custom package that is my own package defined in packagesgo.go
		isvalidName, isvalidEmail, isvalidTicketNumer := validateUserInput(userFirstName, userLastName, email, userTickets, remainingTickets)
		// here validate user can't take more than remaining tickets

		if isvalidName && isvalidEmail && isvalidTicketNumer {
			bookTicket(userFirstName, userLastName, userTickets, remainingTickets, conferenceName, email)
			// call function ffirstName
			firstNames := getFirstNames()
			fmt.Printf("the first name of user %v \n", firstNames)

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

func greetUasers() {
	fmt.Printf("Hello welcome to our %v\n", conferenceName)
	fmt.Printf("Our total ticket %v and remaining ticket %v \n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	// blank identifier is used to ignore the value of the variable _
	for _, booking := range booking {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func validateUserInput(userFirstName string, userLastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// validate user input
	isvalidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumer := userTickets > 0 && userTickets <= remainingTickets
	return isvalidName, isvalidEmail, isvalidTicketNumer
}

func getUserInput() (string, string, string, uint) {
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
	return userFirstName, userLastName, email, userTickets
}

func bookTicket(userFirstName string, userLastName string, userTickets uint, remainingTickets uint, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	// create map
	var userData = make(map[string]string)
	userData["firstName"] = userFirstName
	userData["lastName"] = userLastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	booking = append(booking, userData)
	fmt.Println("Your booking is successfull list\n", booking)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email\n", userFirstName, userLastName, userTickets)
	fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)
}

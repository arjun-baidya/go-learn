// A struct (short for "structure") is a collection of data fields with declared data types. Golang has the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types. Each data field in a struct is declared with a known type, which could be a built-in type or another user-defined type.

// Structs are the only way to create concrete user-defined types in Golang. Struct types are declared by composing a fixed set of unique fields. Structs can improve modularity and allow to create and pass complex data structures around the system. You can also consider Structs as a template for creating a data record, like an employee record or an e-commerce product.

//  The declaration starts with the keyword type, then a name for the new struct, and finally the keyword struct. Within the curly brackets, a series of data fields are specified with a name and a type.

package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Golang Conference"

const conferenceTickets uint = 100

var remainingTickets uint = 100
var booking = make([]userData, 0)

// struct is a collection of fields that together represent a data record different data types
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

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
		firstNames = append(firstNames, booking.firstName) // append to firstNames from struc userData
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
	// create struct and append to booking
	var userData = userData{
		firstName:       userFirstName,
		lastName:        userLastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	booking = append(booking, userData)
	fmt.Println("Your booking is successfull list\n", booking)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email\n", userFirstName, userLastName, userTickets)
	fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)
}

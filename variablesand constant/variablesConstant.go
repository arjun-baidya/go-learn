// Variables in go

// variables are used to store values
// like container for values
// give varibales name reference everywhere in the app

package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Golang Conference"
	const conferenceTickets = 100
	var remainingTickets = 100
	fmt.Println("Hello, Welcome to our Booking App")
	fmt.Println("Welcome to our", conferenceName, "let's enjoy")
	fmt.Println("Get your tickets now")
	fmt.Println("Total tickets available:", conferenceTickets)
	fmt.Println("Remaining tickets:", remainingTickets)

	// formatting out-put
	fmt.Printf("%s %d %d\n", "Total tickets available:", conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to our %v let's enjoy go\n", conferenceName)
	fmt.Printf("Welcome to our %s let's enjoy go\n", conferenceName)
	fmt.Printf("Our total ticket %v and remaining ticket %v", conferenceTickets, remainingTickets)

}

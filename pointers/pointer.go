// a pointer is a variable that holds the memory address of another variable

package main

import "fmt"

func main() {
	var conferenceTickets = 100
	// & is used to get the memory address of the variable
	fmt.Println("address of conferenceTickets:", &conferenceTickets)

}

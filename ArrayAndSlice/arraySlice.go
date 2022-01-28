// Go programming language provides a data structure called the array
// fixed-size sequential collection of elements of the same type
// in go array store same type of data
// In Go, there are two ways to declare an array:
// 1. With the var keyword:
// var array_name = [length]datatype{values}
// 2. With the := sign:
// array_name := [length]datatype{values}
// Two-Dimensional Arrays
// In Go, two-dimensional arrays are declared using the following syntax:
// var array_name = [length1(row)][length2(column)]datatype{values}

// slice in go used for remove limitation array
// array is fixed size but slice size is not fixed when we use slice we can change the size of array
// slice is a reference type and it is not fixed size but it is reference type so we can change the size of slice

package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Golang Conference"
	const conferenceTickets uint = 100
	var remainingTickets uint = 100
	var bookings [100]string
	// slice
	booking := []string{}

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
	bookings[0] = userFirstName + " " + userLastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value of array %v \n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings))
	// slice
	booking = append(booking, userFirstName+" "+userLastName)
	fmt.Printf("The whole slice: %v\n", booking)
	fmt.Printf("The first value of slice %v \n", booking[0])
	fmt.Printf("Slice type %T\n", booking)
	fmt.Printf("slice length %v\n", len(booking))

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you will receieve an confirmation email by %v\n", userFirstName, userLastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v \n", remainingTickets, conferenceName)
}

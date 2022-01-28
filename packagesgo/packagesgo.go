// A package is nothing but a directory with some code files, which exposes different variables (features) from a single point of reference
// A package is a collection of Go source files that are compiled together into a single binary file.

// first declare the package name or folder name then write your code
// your own package name should be unique and should not be same as other package name in your system
// package name should be lower case
// package name should be short and descriptive
// package name should not contain spaces
// package name should not contain special characters
// package name should not start with number
// package name should not start with underscore
// package name should not start with capital letter
// package name should not start with hyphen
// package name should not start with double hyphen
// package name should not start with dot
// package name should not start with plus

// eported function should be in capital letter
package packagesgo

import "strings"

func ValidateUserInput(userFirstName string, userLastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// validate user input
	isvalidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isvalidTicketNumer := userTickets > 0 && userTickets <= remainingTickets
	return isvalidName, isvalidEmail, isvalidTicketNumer
}

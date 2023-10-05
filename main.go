package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference" // This syntax only works for variables without explicit type declaration
const conferenceTickets int = 50

var remainingTickets uint = 50 // uint means that it's a positive int and never goes negative
var bookings = []string{}

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) < 50 {

		// func call for user input
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//call func for booking
			bookings = bookTicket(userTickets, firstName, lastName, email)
			// call function print first names
			fmt.Printf("These are all the bookings: %v\n", getFirstNames())

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}

			if !isValidEmail {
				fmt.Println("Email address is not valid.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to our %v application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and remaining of %v tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Buy your tickets here")
}

func getFirstNames() []string {
	var firstNames = []string{}
	for _, bookings := range bookings { // use _ because we are telling Go that we are not going be using the index variable
		var names = strings.Fields(bookings)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []string {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName) // slices instead of array, slices are an abstraction of an array

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}

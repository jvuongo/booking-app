package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" // This syntax only works for variables without explicit type declaration
const conferenceTickets int = 50

var remainingTickets uint = 50 // uint means that it's a positive int and never goes negative
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// func call for user input
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		//call func for booking
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email) // "go" abstracts the thread creation
		// call function print first names
		fmt.Printf("These are all the bookings: %v\n", getFirstNames())

		var noTicketsRemaining bool = remainingTickets == 0

		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
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
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to our %v application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and remaining of %v tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Buy your tickets here")
}

func getFirstNames() []string {
	var firstNames = []string{}
	for _, bookings := range bookings { // use _ because we are telling Go that we are not going be using the index variable
		firstNames = append(firstNames, bookings.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // slices instead of array, slices are an abstraction of an array
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // stopping this thread for 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################")

	wg.Done()
}

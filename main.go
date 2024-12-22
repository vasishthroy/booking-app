package main

import (
	"booking-app/helper"
	"fmt"
)

// Declaring variables outside the main function
// makes the variable accessible across the script
const conferenceTickets = 50

var conferenceName = "Go conference"

// Making the first letter capital converts the variable into a global variable
// Global variables are accessible across the module or different scripts
var remainingTickets uint = conferenceTickets

// Using array to store the data of users booking the ticket

// Declare an empty array of size 50 of string type
// var bookings = [50]string{}

// Declaring an array with prefilled values
// var bookings = [50]string{"VR", "AP", "TT"}

// Alternate syntax for declaring Array
// var bookings [conferenceTickets]string

// If the size is kept blank then it becomes a Slice
// They are more efficient than arrays as they automatically allocate memory
// based on the amount of value being stored
// Example of simple slice of string dtype
// var bookings []string
// TODO add Descriptive comments for MAPs
var bookings = make([]userDetails, 0)

type userDetails struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	// A for loop executes until the condition for the loop is satisfied
	// eg. for userTickets <= 0 {}
	// Infinite for loop can be implemented in two ways
	// 1) for true {}
	// 2) for {}
	for {

		// Get user details and ticket count
		firstName, lastName, emailAddress, userTickets := getUserInput()

		// Check for any invalid details provided by user
		var isValidName,
			isValidEmail,
			isValidTicketNumber bool = helper.ValidateInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(firstName, lastName, emailAddress, userTickets)

			sendEmail(firstName, lastName, emailAddress, userTickets)

			var firstNames []string
			firstNames = getFirstNames(firstNames)
			fmt.Printf("This is all our current booking: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Thanks you have booked the all our remaining tickets.")
				fmt.Println("Enjoy the show :)")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("Invalid first and last name\n")
			}

			if !isValidEmail {
				fmt.Printf("Invalid email format\n")
			}

			if !isValidTicketNumber {
				fmt.Printf("We do not have %v ticket. ", userTickets)
				fmt.Printf("Remaining Ticket Count: %v\n", remainingTickets)
			}

			continue
		}
	}
}

func greetUser() {
	/**
		This function is used ot print the greeting and information message
		It doesn't take any argument
		It doesn't return any value.
	**/

	// Println ends with a \n
	fmt.Println("Welcome to our", conferenceName, "booking application")
	// "Printf" is Print with format option for adding variables
	// to display variables with the print text
	fmt.Printf("We have a total of %v tickets\n", conferenceTickets)
	fmt.Println("Get your tickets here to attend")
}

// Functions in go can return multiple values
// Although for returning multiple function it needs to be defined
// SYNTAX: func functionName() (return datatype1, return datatype2, ...) {}
func getUserInput() (string, string, string, int) {
	/*
		Function is used to get data from the user like
		First name, last name, email address and no. of tickets the user want to book
		This function doesn't take any argument and returns 4 values
		Return type: string, string, string and int
	*/

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int
	// Ask user their name
	// Scan is used to get input, it takes a pointer as an argument
	// It uses that pointer to assign value to the pointed variable
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email: ")
	fmt.Scan(&emailAddress)

	fmt.Print("Enter the number of tickets you need to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTickets(firstName string, lastName string, emailId string, userTickets int) {
	/*
		This function is used to log the booking done by the user.
		It subtracts the tickets booked by the user from the remaining ticket count
		and creates a log by saving the name of the user inside the booking
		Arguments:
			firstName (string): First Name of the user
			lastName (string): Last Name of the user
			emailId (string): Email of the user
			userTickets (int): No. of tickets the user wants to book.
		This function doesn't return any value.
	*/

	remainingTickets = remainingTickets - uint(userTickets)

	// var userData = userDetails{
	// 	firstName:       firstName,
	// 	lastName:        lastName,
	// 	email:           emailId,
	// 	numberOfTickets: uint(userTickets),
	// }
	var userData userDetails
	userData.firstName = firstName
	userData.lastName = lastName
	userData.email = emailId
	userData.numberOfTickets = uint(userTickets)

	// bookings[0] = firstName + " " + lastName
	// Slice uses the built-in append function to append data to it
	bookings = append(bookings, userData)

	// Check for the array or slice data logged
	/*
		fmt.Printf("Array / Slice: %v\n", bookings)
		fmt.Printf("Array / Slice first element: %v\n", bookings[0])
		fmt.Printf("Array / Slice type: %T\n", bookings)
		fmt.Printf("Array / Slice length: %v\n", len(bookings))
	*/

	// Ask for users name and no. of tickets
	fmt.Printf("Thank you %v %v for booking %v tickets. ",
		firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation on %v soon.\n", emailId)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
}

func getFirstNames(firstNames []string) []string {
	// func getFirstNames(firstNames []string) []string {

	/*
		This function splits the names present inside the booking slice
		into first name and last name, and then saves the fist name
		into the slice argument which is then returned.
		Arguments:
			firstNames (slice string): An empty slice variable
		Return type: slice string
	*/

	// Iterative for loop
	// range lets us iterate over elements
	// for slices and array it provides index and element
	// Go interprets "_" (blank identifier)  as an unused variable
	for _, booking := range bookings {

		// // stings module has a method called Fields
		// // It splits the string based on white spaces available
		// var firstName = strings.Fields(booking)

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func sendEmail(firstName string, lastName string, emailId string, userTicket int) {
	var ticket = fmt.Sprintf("Ticket for %v %v\nTotal Tickets: %v", firstName, lastName, userTicket)
	fmt.Printf("Email sent to %v\n", emailId)
	fmt.Println("####################################")
	fmt.Printf("%v\n", ticket)
	fmt.Println("####################################")

}

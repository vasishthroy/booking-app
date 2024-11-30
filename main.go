package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	// Using array to store the data of users booking the ticket
	// Declare an empty array of size 50 of string type
	// var bookings = [50]string{}
	// Declaring an array with prefilled values
	// var bookings = [50]string{"VR", "AP", "TT"}
	var bookings [conferenceTickets]string

	// Println ends with a \n
	fmt.Println("Welcome to our", conferenceName, "booking application")
	// "Printf" is Print with format option for adding variables
	// to display variables with the print text
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - uint(userTickets)
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Array: %v\n", bookings)
	fmt.Printf("Array first element: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	// Ask for users name and no. of tickets
	fmt.Printf("Thank you %v %v for booking %v tickets. ",
		firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation on %v soon.\n",
		emailAddress)
	fmt.Printf("Remaining tickets: %v", remainingTickets)

}

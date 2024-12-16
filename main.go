package main

import (
	"fmt"
	"strings"
)

func main() {
	// Alternative Syntax of creating a variable with predefined value
	conferenceName := "Go conference"
	const conferenceTickets = 50
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
	var bookings []string

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// A for loop executes until the condition for the loop is satisfied
	// eg. for userTickets <= 0 {}
	// Infinite for loop can be implemented in two ways
	// 1) for true {}
	// 2) for {}
	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(emailAddress, "@")
		isValidTicketNumber := userTickets > 0 && remainingTickets >= uint(userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - uint(userTickets)
			// bookings[0] = firstName + " " + lastName
			// Slice uses the built-in append function to append data to it
			bookings = append(bookings, firstName+" "+lastName)

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
			fmt.Printf("You will receive a confirmation on %v soon.\n",
				emailAddress)
			fmt.Printf("Remaining tickets: %v\n", remainingTickets)

			var firstNames []string

			// Iterative for loop
			// range lets us iterate over elements
			// for slices and array it provides index and element
			// Go interprets "_" (blank identifier)  as an unused variable
			for _, name := range bookings {

				// stings module has a method called Fields
				// It splits the string based on white spaces available
				var firstName = strings.Fields(name)
				firstNames = append(firstNames, firstName[0])
			}

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

func greetUser(confName string, confTickets int, remainingTickets uint) {
	// Println ends with a \n
	fmt.Println("Welcome to our", confName, "booking application")
	// "Printf" is Print with format option for adding variables
	// to display variables with the print text
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, int) {
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

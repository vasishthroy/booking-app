package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	// Println ends with a \n
	fmt.Println("Welcome to our", conferenceName, "booking application")
	// "Printf" is Print with format option for adding variables
	// to display variables with the print text
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// Ask user their name
	// Scan is used to get input, it takes a pointer as an argument
	// It uses that pointer to assign value to the pointed variable
	fmt.Print("Please enter your name: ")
	fmt.Scan(&userName)

	// Ask for users name and no. of tickets
	fmt.Printf("%v booked %v tickets.", userName, userTickets)

}

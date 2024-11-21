package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// Ask for users name and no. of tickets
	fmt.Printf("%v booked %v tickets.", userName, userTickets)

}

package helper

import "strings"

// Keeping the first letter of the function capital makes it a global function
// Global function can be imported across different scripts
func ValidateInput(firstName string, lastName string, emailId string, ticketNeeded int, remainingTickets uint) (bool, bool, bool) {
	/*
		Check if user has not provided invalid data
		Arguments:
			firstName (string): First name of the user.
			lastName (string): Last name of the user.
			emailId (string): Email id of the user.
			ticketNeeded (int): No. of ticket user needs.
		Return type: bool, bool and bool

	*/

	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(emailId, "@")
	isValidTicketNumber := ticketNeeded > 0 && remainingTickets >= uint(ticketNeeded)
	return isValidName, isValidEmail, isValidTicketNumber
}

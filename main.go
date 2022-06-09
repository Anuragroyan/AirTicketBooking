package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 1)

// waitgp for maintaining thread execution
var waitgp = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, emailID, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailID, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, emailID)
		waitgp.Add(1)
		// concurrency using go keyword
		go sendTicket(userTickets, firstName, lastName, emailID)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {

			fmt.Println("Our conference is booked out. Come back next year. ")

		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you enterted is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	waitgp.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// logic for extract firstname from user
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, emailID string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailID, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// logic for taking user input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailID string
	var userTickets uint

	fmt.Println("Enter your FirstName: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your LastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailID)

	fmt.Println("Enter number of tickets to be booked: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailID, userTickets
}

// logic for map using to store userdata in one container
func bookTicket(userTickets uint, firstName string, lastName string, emailID string) {
	remainingTickets = remainingTickets - userTickets
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailID"] = emailID
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n\n", bookings)
	fmt.Printf(" (',') Thank you %v %v for booking %v tickets. You will receive a confirmation on your email address %v.\n\n", firstName, lastName, userTickets, emailID)
	fmt.Printf("No of %v tickets are  available for booking for %v.\n\n", remainingTickets, conferenceName)
}

// logic for sending ticket to emailid
func sendTicket(userTickets uint, firstName string, lastName string, emailID string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, emailID)
	fmt.Println("###################")
	waitgp.Done()
}

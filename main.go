package main

import (
	"fmt"
	"time"
	"sync"
)


const conferenceTickets int = 20
var remainingTickets uint = 20
var conferenceName = "GO Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	tickets uint
}

var wg = sync.WaitGroup{}


func main() {

	greetUsers()

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendticket(userTickets, firstName, lastName, email)
			
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
				
				if remainingTickets == 0 {
					fmt.Println("Our conference is booked out. Come back next year!")
					
				}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is not valid. It's too short")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid. It doesn't contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
		wg.Wait()
		
	}
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
	firstNames = append(firstNames, booking.firstName)
		}	
	return firstNames
}

func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket (userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		tickets: userTickets,
	}


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendticket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################")
	wg.Done()
}
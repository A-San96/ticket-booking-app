package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked for %v tickets.\n", userName, userTickets)

}

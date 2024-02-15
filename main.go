package main

import "fmt"
import "strings"
var conferenceName string
var remTickets uint
var bookings []string
func main() {
	conferenceName = "Go Conference"
	const TOTAL_TICKETS = 50
	remTickets = TOTAL_TICKETS
	greeting()

	bookings = []string{}


	for remTickets > 0{
		//name and number
		var fName string
		var lName string
		var userTickets uint
		var email string
		fmt.Println("Enter your first name:")
		fmt.Scan(&fName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("How many tickets would you like?")
		fmt.Scan(&userTickets)

		if len(fName) < 2 || len(lName) < 2 {
			fmt.Printf("invalid name %v", fName + " " + lName)
			continue
		}
		if strings.Contains(email, "@") {
			fmt.Printf("invalid email %v", email)
			continue
		}
		if userTickets > remTickets{
			fmt.Printf("We don't have %v tickets\n", userTickets)
			continue
		}
		fmt.Printf("Thanks %v %v for booking %v tickets\n", fName, lName, userTickets )
		remTickets -= userTickets
		bookings = append(bookings, fName + " " + lName)

		fmt.Printf("full array %v\n", bookings)
		fmt.Printf("%v tickets remaining\n", remTickets)

		
		fmt.Printf("All bookings %v\n", retFNames())

		if remTickets == 0 {
			fmt.Print("sold out\n")
			break;
		}

	}

	
}


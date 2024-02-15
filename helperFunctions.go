package main

import (
	"fmt"
	"strings"
)

func retFNames() []string {
	firstNames := []string{}

	for _, booking := range bookings{
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}
func greeting() {
	fmt.Printf("Welcome to the %v ticketing application\n", conferenceName)
	fmt.Printf( "%v are remaining\n", remTickets)
}
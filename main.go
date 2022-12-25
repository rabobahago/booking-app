package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"
    const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	//array
	//var booking = [50]string{}
	//var bookings [50]string
	//slice
	for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Printf("Please Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Please Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Please Enter your email address: ")
		fmt.Scan(&email)
		fmt.Printf("Please Enter number of tickets: ")
		fmt.Scan(&userTickets)
		var isValidName bool = len(firstName) >= 2  && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValidEmail && isValidTicketNumber && isValidName {
			remainingTickets = remainingTickets - userTickets
		//adding element to array
		//bookings[0] = firstName + " " + lastName

		bookings = append(bookings, firstName + " " + lastName)
		// fmt.Printf("The whole slice %v\n", bookings)
		// fmt.Printf("The first value of slice %v\n", bookings[0])
		// fmt.Printf("The type of an slice %T\n", bookings)
		// fmt.Printf("The size of an slice %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking);
			var firstName = names[0];
           firstNames = append(firstNames, firstName)
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		var noTicketRemaining bool = remainingTickets == 0
		if noTicketRemaining {
			fmt.Printf("Our conference is booked out. Please come back next year")
			break
		}
		}else{
			fmt.Printf("your input data is invalid, please try again\n")

		}


	}

}
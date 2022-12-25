package main

import (
	"fmt"
	"strconv"
)
var conferenceName = "Go Conference"
    const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings = make([]map[string]string, 0)

func main(){


	greetUsers()
	//array
	//var booking = [50]string{}
	//var bookings [50]string
	//slice
	for{
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := validateUser(firstName, lastName, email, userTickets, remainingTickets)
		if isValidEmail && isValidTicketNumber && isValidName {
			remainingTickets  :=bookingTicket(remainingTickets, userTickets, firstName, lastName, conferenceName,email)
		firstNames :=  getFirstName()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
		var noTicketRemaining bool = remainingTickets == 0
		if noTicketRemaining {
			fmt.Printf("Our conference is booked out. Please come back next year")
			break
		}
		}else{
			if !isValidEmail{
				fmt.Printf("The email you entered doesn't contain @ symbol\n")
			}
			if !isValidName{
				fmt.Printf("The name you entered is too short\n")
			}
			if !isValidTicketNumber{
				fmt.Printf("The number you entered is not valid\n")
			}
			fmt.Printf("your input data is invalid, please try again\n")

		}


	}

}
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstName() []string {
	firstNames := []string{}

		for _, booking := range bookings {
           firstNames = append(firstNames, booking["firstName"])
		}
		return firstNames
}

func getUserInput() (string, string,string, uint){
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
		return firstName, lastName, email, userTickets
}
func bookingTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, conferenceName string, email string) uint{
	    remainingTickets = remainingTickets - userTickets
		//adding element to array
		//bookings[0] = firstName + " " + lastName
        var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v\n", bookings)
		// fmt.Printf("The whole slice %v\n", bookings)
		// fmt.Printf("The first value of slice %v\n", bookings[0])
		// fmt.Printf("The type of an slice %T\n", bookings)
		// fmt.Printf("The size of an slice %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
		return remainingTickets
}
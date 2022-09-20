package main 
import "fmt"

func main(){
	conferenceName := "Go Conference"
	bookings := []string{}
	var usrName string
	var remainingTickets uint = 50
	var reqTickets uint
	var email string
	for {
	  fmt.Printf("Welcome to %v\n",conferenceName)
	  fmt.Printf("We have total of %v tickets\n ",remainingTickets)
	  //fmt.Printf("%T",conferenceName)
	  fmt.Println("Please enter your name")
	  fmt.Scan(&usrName)
	  fmt.Println("Please enter your email")
	  fmt.Scan(&email)
	  fmt.Println("Please enter the number of tickets you want")
	  fmt.Scan(&reqTickets)
	  remainingTickets = remainingTickets - reqTickets
	  bookings = append(bookings, usrName)
	 /* fmt.Printf("%v\n",bookings)
	  fmt.Printf("%v\n",bookings[0])
	  fmt.Printf("%T\n",bookings)
	  fmt.Printf("%v\n",len(bookings)) 
	 */
	 fmt.Printf("Out of 50 tickets we have sold %v tickets and still left with %v tickets",reqTickets,remainingTickets)
    }
}

// isValidName := firstName >= 2 && lastName >=2
// isValidEmail := strings.Contains(email,"@")
// isValidTicketNumer := userTickets > 0 && userTickets <= remainingTickets
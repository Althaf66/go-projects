package main

import (
  "fmt"
  "app/helper"
  "time"
  )
  var name string = "go conference"
	const conferenceticket = 50
	var remaining int = 50
  var booking = make([]Userdata, 0)

  type Userdata struct{
    firstname string
    lastname string
    email string
    ticket int
  }

func main(){
  
  greetuser(name,conferenceticket)
  for{
    userf,userl,ticket,email := getUserInput()

  isValidName,isValidEmail,isValidTickets :=  helper.ValidateInput(userf,userl,email,ticket,remaining)
    
    if isValidName && isValidEmail && isValidTickets{
      BookTicket(ticket,userf,userl,email)
      go sendticket(ticket,userf,userl,email)
      
      firstnames := getFirstname()
      fmt.Printf("first names of bookings:%v\n",firstnames)

    if remaining == 0 {
      fmt.Println("our conference is booked out.come back next year")
      break
    }
    }else{
      if !isValidName {
         fmt.Println("first name or last name is too short")
      }
      if !isValidEmail {
        fmt.Println("this input does not contain @ symbol")
      }
      if !isValidTickets {
        fmt.Println("number of tickets you have entered is invalid")
      }
    }
  }
}

//function to greet user
func greetuser(confname string,ticketnum int){
  fmt.Printf("welcome to the %v booking app\n",confname)
  fmt.Println("Get your ticket")
	fmt.Printf("We have total tickets of %v tickets\n",ticketnum)
}

// function to print first name only in the array
func getFirstname() []string{
  firstn := []string{}
    for _,booking := range booking {
      firstn = append(firstn, booking.firstname)
    }
  return firstn
}

//function to enter user's information
func getUserInput()(string,string,int,string){
  var userf string
	var userl string
	var ticket int
	var email string

  fmt.Println("enter your first name")
	fmt.Scan(&userf)
	fmt.Println("enter your last name")
	fmt.Scan(&userl)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&ticket)
	fmt.Println("Enter your email account")
	fmt.Scan(&email)

  return userf,userl,ticket,email
}

func BookTicket(ticket int,userf string,userl string,email string){
  remaining = conferenceticket - ticket
  
  var userData = Userdata {
    firstname: userf,
    lastname: userl,
    email: email,
    ticket: ticket,
  }
  booking = append(booking,userData)
  fmt.Printf("List of bookings is %v\n",booking)
	fmt.Printf("Thank you, %v %v for booking %v tickets.Confirmation were sent to %v\n",userf,userl,ticket,email)
  fmt.Printf("remaining of %v tickets",remaining)
}

func sendticket(ticket int,userf string,userl string,email string){
  time.Sleep(3*time.Second)
  var tickets = fmt.Sprintf("%v ticket for %v %v.",ticket,userf,userl)
  fmt.Println("\n#############################\n")
  fmt.Printf("Sending ticket.....\n %v message sended to email address %v",tickets,email)
  fmt.Println("\n#############################\n")
}
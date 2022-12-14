package main

import "fmt"

func main() {
	var firstName string
	var lastName string

	fmt.Print("Enter Friend First Name : \n")
	fmt.Scanln(&firstName)
	fmt.Print("Enter Friend Last Name : \n")
	fmt.Scanln(&lastName)

	fmt.Print("Your Friend Full Name is : ")
	fmt.Print(firstName, " "+lastName)
}

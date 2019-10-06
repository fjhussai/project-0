package main

import "fmt"

func main() {
	var grting [3]string
	grting[0] = "Hi there! Welcome to Piggy Bank! \n"
	grting[1] = "Piggy Bank is here to help you keep track of your savings."
	grting[2] = "Let's get started!"

	fmt.Println(grting[0], grting[1])
}

//Have users set a username
/*
func getuser() {
	fmt.Println("Please enter a username below:")
	var username string
	fmt.Scanln(&username)
}

type users struct {
	username string
	password string
	email    string
}


//have users set a password
//Have the pairs stored in a map
//Keep track of user data in a structure

*/

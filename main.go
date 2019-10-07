package main

import "fmt"

func main() {
	greeting()
	getuser()
}

func greeting() {
	var grting [3]string
	grting[0] = "Hi there! Welcome to Piggy Bank! \n"
	grting[1] = "Piggy Bank is here to help you keep track of your savings."
	grting[2] = "Let's get started!"

	fmt.Println(grting[0], grting[1])
}

type userinfo struct {
	username string
	password string
	email    string
}

//Have users set a username

func getuser() {
	fmt.Println("Please enter a username below:")
	var collect userinfo
	collect.username = "cardib"
	fmt.Println(collect.username)

	fmt.Println("Please choose a password.")
	collect.password = "bodakyellow"
	fmt.Println(collect.password)

}

//have users set a password

//Have the pairs stored in a map
//Keep track of user data in a structure

package main

import (
	"fmt"
)

func greeter(){
	fmt.Println("Hello! Welcome to PiggyBank. Are you a customer or an employee?")
	var input string
	fmt.Scan(&input)

	if input = "customer" {
//have them enter their login/password
// or have them create an account
	} else if input = "employee" {
		//have the employee login stuff
	} else {
		fmt.Println("Please enter a valid input of either 'employee' or 'customer'.")
	}
}

func authenticate( {
	var givenuniqname string
	var givenpass string
	fmt.Print("uniqname:")
	fmt.Scan(&givenuniqname)
	fmt.Print("password:")
	fmt.Scan(&givenpass)

	sqlquery := `select uniqname from customer where uniqname=&givenuniqname`

})
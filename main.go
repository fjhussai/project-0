package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	ping(db)
	showUserTable(db)
	showAcctsTable(db)
	getacctnum()

}

func ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

//this function prints out all the values in the user accounts tabel except password
func showUserTable(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM user_accounts")
	for rows.Next() {
		var uniqname string
		var userfirst string
		var userlast string
		var password string
		var funds float64

		rows.Scan(&uniqname, &userfirst, &userlast, &password, &funds)
		fmt.Println(uniqname, userfirst, userlast, funds)
	}
}

func showAcctsTable(db *sql.DB) {
	rows, _ := db.Query(`SELECT * FROM bank_accounts`)
	for rows.Next() {
		var acctnumber int
		var uniqname string
		var acctbalance float64
		var accttype string

		rows.Scan(&acctnumber, &uniqname, &acctbalance, &accttype)
		fmt.Println(acctnumber, uniqname, acctbalance, accttype)
	}
}

/* func whattodo() {
	fmt.Println("What would you like to do today?")
	fmt.Println("To check your account balance, choose 1")
	fmt.Println("to make a deposit, choose 2")
	fmt.Println("To make a withdrawal, choose 3")
	fmt.Println("To open a new account, choose 4")
	fmt.Println("To transfer funds between accounts, choose 5")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createacct(db)
	}

	fmt.Println("thank you for using piggy bank!")
}
*/

//deposit is a function that will ask for an account number, get its balance, and then add to it

/*
func deposit() {
	fmt.Println("hello. Please enter an account number")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM bank_accounts WHERE acct_number = $1", searchvalue)
	var acctbalance float64
	row.Scan(&acctbalance)
	fmt.Println(acctbalance)
}
*/

//this is a function that creates user accounts
func createacct(db *sql.DB) {
	var uniqname string
	var userfirst string
	var userlast string
	var password string

	fmt.Println("Please choose a unique username. Note that no one else can have the same username as you.")
	fmt.Scan(&uniqname)

	fmt.Println("Please enter your first name.")
	fmt.Scan(&userfirst)

	fmt.Println("Please enter your last name.")
	fmt.Scan(&userlast)

	fmt.Println("Please choose a password")
	fmt.Scan(&password)

	sqlStatement := `INSERT INTO user_accounts values ($1,$2,$3,$4,$5)`
	db.Exec(sqlStatement, uniqname, userfirst, userlast, password, 0)
}

// generates a pseudo-random number
//find a way to put this into the acountid in bank accounts table

func getacctnum() {
	var num int
	num = rand.Intn(10000)
	fmt.Print(num)

	s := strconv.Itoa(num)

	db.Exec(`update user_accounts set acctnumber=num where uniqname = givenuniqname`)
}

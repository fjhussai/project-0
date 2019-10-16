package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

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
	//userauth(db)
	//createbankacct(db)
	//deposit(db)
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

// this function creates a new bank account
func createbankacct(db *sql.DB) {
	var givenuniqname string
	fmt.Println("Please enter your uniqname")
	fmt.Scan(&givenuniqname)

	acctnumber := getacctnum()
	acctbalance := 0

	var accttype string
	fmt.Println("What kind of account is this? ex. 'checking'")
	fmt.Scan(&accttype)

	sqlStatement := `INSERT INTO bank_accounts values($1,$2,$3,$4)`
	_, err := db.Exec(sqlStatement, acctnumber, givenuniqname, acctbalance, accttype)

	if err != nil {
		fmt.Print("An error occurred creating your account.", err)
	}
}

// generates a pseudo-random number
//find a way to put this into the acountid in bank accounts table

func getacctnum() string {
	rand.Seed(time.Now().UnixNano())
	var num int
	num = rand.Intn(10000)
	fmt.Println(num)

	acctnumstr := strconv.Itoa(num)
	return acctnumstr
}

//deposit is a function that will ask for an account number, get its balance, and then add to it
func deposit(db *sql.DB) {
	fmt.Println("hello. Please enter an account number")
	var searchvalue string
	fmt.Scanln(&searchvalue)

	row := db.QueryRow("SELECT acctbalance FROM bank_accounts WHERE acctnumber = $1", searchvalue)

	var acctbalance float64
	err := row.Scan(&acctbalance)
	fmt.Println("Your previous balance was ", acctbalance)
	if err != nil {
		fmt.Println(err)
	}

	var depositamt float64
	fmt.Println("how much would you like to deposit in this account?")
	fmt.Scan(&depositamt)

	newbalance := depositamt + acctbalance
	db.Exec("UPDATE bank_accounts SET acctbalance = $1 WHERE acctnumber = $2", newbalance, searchvalue)
	fmt.Println("Your new account balance is ", newbalance)
}

/*
//this function looks up an account number and returns an account balance
func getacctbalance() {
	fmt.Println("Please enter an account number")
	var givenacctnum string
	fmt.Scan(&givenacctnum)
	sqlcommand := db.Query(`select acctbalance from bank_accounts where acctnumber = givenacctnum`)
	for sqlcommand.Next() {

	}
}
*/
//Any function to deposit or withdraw should call this

//this is a function to check your login info
func userauth(db *sql.DB) {
	var collectuniqname string
	var collectpassword string

	fmt.Print("uniqname:")
	fmt.Scan(&collectuniqname)

	fmt.Print("password:")
	fmt.Scan(&collectpassword)

	rows, err := db.Query("SELECT * FROM user_accounts WHERE uniqname = $1 AND password = $2", collectuniqname, collectpassword)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var uniqname string
		var userfirst string
		var userlast string
		var password string
		var funds float64

		err = rows.Scan(&uniqname, &userfirst, &userlast, &password, &funds)
		if err != nil {
			fmt.Println("Error scanning row", err)
			continue
		}
		fmt.Println("Welcome,", uniqname)
	}
}

package main

import (
	"encoding/json"
	"fmt"
)

type BankInfo struct {
	User    User   `json:"user"`
	Account string `json:"account"`
}

type User struct {
	Uniqname  string  `json:"uniqname"`
	Userfirst string  `json:"userfirst"`
	Userlast  string  `json:"userlast"`
	Password  string  `json:"password"`
	Funds     float64 `json: "funds"`
}

func main() {
	fmt.Println("Hello World")

	user := User{Uniqname: "cardib", Userfirst: "Cardi", Userlast: "Belcalis", Password: "bodakyellow", Funds: 540}
	bankinfo := BankInfo{User: user, Account: "54"}

	fmt.Printf("%+v\n", bankinfo)

	bytearray, err := json.MarshalIndent(bankinfo, " ", "  ")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(bytearray))

}

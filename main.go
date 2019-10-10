package main

func main() {
	greeting.greeting()
	newuser()
}

//userinfo is a structure to represent user information
type Userinfo struct {
	uniqname  string
	userfirst string
	userlast  string
	password  string
	funds     float32
}

func newuser(uniqname string, userfirst string, userlast string, password string, funds float32) Userinfo {
	return Userinfo{
		uniqname:  uniqname,
		userfirst: userfirst,
		userlast:  userlast,
		password:  password,
		funds:     funds,
	}
}

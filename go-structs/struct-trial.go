package main

import "fmt"

var userData [2]Users
var user1 Users
var user2 Users

func main() {

	user1 = Users{"1", "tiq", "2412"}
	user2 = Users{"2", "tiqdesign", "3461"}
	userData = [2]Users{user1, user2}
	/*
		fmt.Print("1.Kullanıcı;")
		fmt.Println("Kullanıcı Adı:", userData[0].username)
		fmt.Println("Şifre:", userData[0].password, "\n")

		fmt.Print("2.Kullanıcı;")
		fmt.Println("Kullanıcı Adı:", userData[1].username)
		fmt.Println("Şifre:", userData[1].password)*/
	fmt.Println("------------------------------------")
	printUsers()

}

type Users struct {
	userId   string
	username string
	password string
}

func printUsers() {
	for i := 0; i < len(userData); i++ {
		fmt.Println("Kullanıcı Adı:", userData[i].username)
		fmt.Println("Şifre:", userData[i].password)
		fmt.Println("------------------------------------")
	}
}

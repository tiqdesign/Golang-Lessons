package main

import (
	"fmt"
)

func main() {
	arrayExample()
	slicesExample()
	mapExample()
}

func arrayExample() {
	//we have to assign the range of the array
	bookArray := [2]string{"Satranç", "O muydu?"}
	//we can print one by one or
	book1 := bookArray[0]
	book2 := bookArray[1]
	fmt.Println(book1, "-", book2)

	//we can print foreach item on the bookArray
	for _, item := range bookArray {
		fmt.Println(item)
	}
}

func slicesExample() {
	userSlices := make([]Users, 3)
	user1 := Users{ID: 1, Username: "Tarık", Password: "asd"}
	user2 := Users{ID: 2, Username: "tiqdesign", Password: "2412"}
	user3 := Users{ID: 3, Username: "tiq24", Password: "tiqdesign"}

	userSlices[0] = user1
	userSlices[1] = user2
	userSlices[2] = user3

	/*addItems := append(userSlices, user3, user2, user1)
	fmt.Println(addItems)*/

	for i := 0; i < len(userSlices); i++ {
		fmt.Println(userSlices[i])
	}
}

type Users struct {
	ID       int
	Username string
	Password string
}

func mapExample() {
	mapForUser := make(map[Users]int)
	user4 := Users{ID: 4, Username: "Tiqdeveloping", Password: "2412"}
	user5 := Users{ID: 5, Username: "Tiqdeveloping", Password: "2412"}

	mapForUser[user4] = 4
	mapForUser[user5] = 5
	fmt.Println(mapForUser[user4], mapForUser[user5])
}

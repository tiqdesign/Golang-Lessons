package main

import (
	"fmt"
	"math/rand"
	"time"
)

var st string

func main() {

	//Random bir sayı üretip tek mi çift mi diye kontrol ediyoruz.
	//Zamana bağlı olarak sürekli yeni bir sayı üretiyor
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n1 := r1.Intn(100)

	//if else yazarken else ten önce parantezin kapandıgından ve else in solunda oldugundan emin ol!!
	/*if number%2 == 0 {
		fmt.Println(number, "is even number")
	} else {
		fmt.Println(number, "is odd number")
	}*/
	stuation := oddOrEven(n1 % 2)
	fmt.Println(n1, "is", stuation)
}

func oddOrEven(num int) string {

	switch num {
	case 0:
		st = "an Even Number"
		break
	case 1:
		st = "an Odd Number"
		break
	}
	return st
}

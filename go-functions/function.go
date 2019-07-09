package main

import "fmt"

func main() {

	//variadic function = değişken sayıda argüman alabilen fonksiyondur.
	//Burada biz bir array yolladık ama bunun dışında bu arrayın istediğimiz bir kısmını kesip yollayabilirdik.
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	variadicFunc(number...)

	example1 := add(2, 8)
	fmt.Println(example1)
	example2 := multiply(2, 8)
	fmt.Println(example2)
	example3 := divide(2, 0)
	fmt.Println(example3)
	//-----------------------------
	num, name := multireturn()
	fmt.Println(name, num)
}

func add(val1, val2 int) int {
	result := val1 + val2
	return result
}
func delete(val1, val2 int) int {
	result := val1 - val2
	return result
}
func multiply(val1, val2 int) int {
	result := val1 * val2
	return result
}

//burada sonuc float olmalı düzelt
func divide(val1, val2 int) int {
	if val2 == 0 {
		fmt.Println("0 a bölünemez!")
		return -1
	} else {
		result := val1 / val2
		return result
	}
}
func multireturn() (int, string) {
	return 24, "Tarık"
}
func variadicFunc(number ...int) {
	sum := 0
	for _, num := range number {
		sum += num
	}
	fmt.Println(sum)
}

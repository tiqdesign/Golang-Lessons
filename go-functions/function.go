package main

import "fmt"

func main() {

	//variadic function = değişken sayıda argüman alabilen fonksiyondur.
	//Burada biz bir array yolladık ama bunun dışında bu arrayın istediğimiz bir kısmını kesip yollayabilirdik.
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	variadicFunc(number...)
	//----------------------------
	example1 := add(2, 8)
	fmt.Println(example1)
	example2 := multiply(2, 8)
	fmt.Println(example2)
	example3 := divide(2, 0)
	fmt.Println(example3)
	//-----------------------------
	num, name := multireturn()
	fmt.Println(name, num)
	//-----------------------------
	nextInt := intSeq()
	//nextInt i sanki bir fonksiyonmuş gibi yazıyoruz nedeni ise içerisindeki isimsiz fonksiyonu çalıştırmak.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//bir daha fonksiyonu baska bir isimle calıstırıp bir değere atarsak bu sefer en bastan başlar
	newInts := intSeq()
	fmt.Println(newInts())

	//output : 120 = 5!
	fact := recursionFunc(5)
	fmt.Println(fact)
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

//Closure - içerisinde anonymous function bulundurur( herhangi bir ismi olmayan func)
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Recursion Function = kendi kendini cagıran fonksiyonlardır. Faktöriyel hesabı yaparak denemiş olduk
func recursionFunc(number int) int {

	if number == 0 {
		return 1
	} else {
		return number * recursionFunc(number-1)
	}
}

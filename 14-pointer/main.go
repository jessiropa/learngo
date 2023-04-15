package main

import "fmt"

func main() {

	// 1. Penerapan Pointer

	var numberA int = 4
	var numberB *int = &numberA // pointer

	fmt.Println("numberA (Value)   :", numberA)
	fmt.Println("numberA (Address) :", &numberA)

	fmt.Println("numberB (Value)   :", *numberB)
	fmt.Println("numberB (Address) :", numberB)

	// 2. Efek perubahan nilai pointer
	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (Value)   :", numberA)
	fmt.Println("numberA (Address) :", &numberA)

	fmt.Println("numberB (Value)   :", *numberB)
	fmt.Println("numberB (Address) :", numberB)

	// 3. Parameter pointer 
	var num = 4
	fmt.Println("before :", num)

	change(&num, 10)
	fmt.Println("after :", num)
}

// 3. Parameter pointer 
func change(original *int, value int){
	*original = value
}
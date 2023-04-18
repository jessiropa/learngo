package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. penerapan keyword defer

	// defer fmt.Println("Hello")
	// fmt.Println("Selamat datang")

	// 2. defer menggunakan fungsi
	// orderSomeFood("pizza")
	// orderSomeFood("burger")

	// number := 3

	// 3. percabangan biasa
	// if number == 3 {
	// 	fmt.Println("haaaai 1")
	// 	defer fmt.Println("haaaai 3")
	// }

	// 3. Percabangan menggunakan IIFE
	// if number == 3 {
	// 	fmt.Println("haaaai 1")
	// 	func(){
	// 		defer fmt.Println("haaaai 3")
	// 	}()
	// }

	// fmt.Println("haaaai 2")

	// 4. Penerapan fungsi os.exit
	defer fmt.Println("hallo")
	os.Exit(1)
	fmt.Println("selamat datang")
}

// func orderSomeFood(menu string){
// 	defer fmt.Println("Terima kasih, silahkan tunggu")
// 	if menu == "pizza"{
// 		fmt.Print("Pilihan tepat!", " ")
// 		fmt.Print("Pizza ditempat kami paling enak!", "\n")
// 		return
// 	}

// 	fmt.Println("Pesanan anda: ", menu)
// }
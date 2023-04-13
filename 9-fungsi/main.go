package main

import (
	"fmt"
	// "strings"
	"math/rand"
	"time"
)

var randomnizer = rand.New(rand.NewSource(time.Now().Unix()))
func main(){
	fmt.Println("BELAJAR FUNGSI")
	
	// Menggabungkan string dalam bentuk array
	// var names = []string{"Jessi", "Kristin"}
	// printMessage("Hallo", names)

	// menampilkan nilai random yang ditentukan

	// var nilaiRandom int 

	// memanggil fungsi nilai random 
	// nilaiRandom = rangeRandom(3,9)
	// fmt.Println("Nilai random : ", nilaiRandom)
	// nilaiRandom = rangeRandom(3,9)
	// fmt.Println("Nilai random : ", nilaiRandom)
	// nilaiRandom = rangeRandom(3,9)
	// fmt.Println("Nilai random : ", nilaiRandom)

	// memanggil fungsi pengecekan pembagian 
	bagiAngka(10, 2)
	bagiAngka(9, 3)
	bagiAngka(4, 2)
	bagiAngka(4, 0)
}

// func printMessage(message string, arr []string){
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// fungsi untuk menampilkan nilai random berdasarkan range 

// func rangeRandom (min, max int) int {
// 	var value = randomnizer.Int()%(max-min+1) + min
// 	return value
// }

// fungsi pengecekan 

func bagiAngka(a, b int){
	if b == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", a, b)
		return
	}

	var res = a / b 
	fmt.Printf("%d/%d = %d \n", a, b, res)
}
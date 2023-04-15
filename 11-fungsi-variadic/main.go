package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Belajar Fungsi Variadic")
	// jumlah parameter pada fungsi variadic bisa banyak parameter
	var avg = calculate(2,4,3,5,4,3,3,5,5,3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pemanggilan hobbies
	yourHobbies("Jessi", "Makan", "tidur", "belajar", "terus berulang")
}

// fungsi variandic 
// nomor merupakan parameter variadic
func calculate(nomor ...int) float64{
	var total int = 0
	for _, angka := range nomor{
		total += angka
	}
	var avg = float64(total) / float64(len(nomor))
	return avg
}

// fungsi dengan parameter biasa dan variadic

func yourHobbies(name string, hobbies ...string){
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s \n", name)
	fmt.Printf("My hobbies are: %s \n", hobbiesAsString)
}
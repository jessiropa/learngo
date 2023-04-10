package main

import "fmt"

func main(){
	fmt.Println("BELAJAR OPERATOR DI GOLANG!!!")

	// Operator aritmatika 
	nilai := (((2 + 6) % 3) * 4-2)/3
	nilai1 := (3 + 5) / (2+2)
	nilai2 := 3 + 10 / 2 + 2
	fmt.Println(nilai)
	fmt.Println(nilai1)
	fmt.Println(nilai2)

	// operator perbandingan
	cek := (nilai == 2)
	fmt.Println(cek)

	// operator logika
	benar := true
	salah := false

	operatorDan := benar && salah
	fmt.Println(operatorDan)

	operatorAtau := benar || salah
	fmt.Println(operatorAtau)

	operatorNot := !salah
	fmt.Println(operatorNot);
}
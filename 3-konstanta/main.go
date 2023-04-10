package main

import "fmt"

func main(){
	fmt.Println("BELAJAR KONSTANTA GOLANG!!!")

	// deklarasi konstanta tunggal 
	const phi float32 = 3.14
	fmt.Print(phi)

	const planet = "mars"
	fmt.Print("\nNama planet favorit", planet) //tanpa spasi di akhir kalimat
	fmt.Print("\nNama planet favorit ", planet) // menggunakan spasi di akhir kalimat

	// deklarasi konstanta multi konstanta
	// contoh 1
	const (
		bullet string= "bulat"
		red bool = true
		angka uint = 9
		desimal float32 = 3.14
	)
	fmt.Println("\nMulti konstanta", bullet, red, angka, desimal)
	fmt.Printf("Tipe data bullet %T \n", bullet)
	fmt.Printf("Tipe data red %T \n", red)
	fmt.Printf("Tipe data angka %T \n", angka)
	fmt.Printf("Tipe data desimal %T \n", desimal)

	// deklarasi konstanta dengan tipe yang sama
	const (
		a = 1
		b 
	)
	fmt.Printf("Tipe data a %T \n", a)
	fmt.Printf("Tipe data a %T \n", b)

	// deklarasi multiple konstanta dalam satu baris 
	const aku, kamu = "watashi", "anata"
	const one, two = 1, 2
	fmt.Println("Nilai aku = ", aku)
	fmt.Printf("Tipe data aku %T \n", aku)
	fmt.Println("Nilai kamu = ", kamu)
	fmt.Printf("Tipe data kamu %T \n", kamu)
	fmt.Println("Nilai one = ", one)
	fmt.Printf("Tipe data one %T \n", one)
	fmt.Println("Nilai two = ", two)
	fmt.Printf("Tipe data two %T \n", two)
}
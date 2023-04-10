package main

import "fmt"

func main() {
	// belajar komentar
	// ini komentar tunggal
	/*
		ini komentar multiline
	*/
	// Belajar variabel
	// deklarasi variabel menggunakan konsep manyfest typing
	// println("Deklarasi variabel dengan keyword var / manyfest typing")
	// var firstName string = "Jessi"
	// fmt.Println("Hallo", firstName)

	// deklarasi variabel menggunakan konsep type inference
	// println("Deklarasi variabel dengan konsep type inference")
	// waktu := "siang"
	// fmt.Println("Selamat", waktu)

	// var suhu string = "dingin"
	// fmt.Println("Cuacanya adalah ", suhu)

	// deklarasi multi variabel
	// var buah1, buah2, buah3 string = "mangga", "pisang", "jeruk"
	// fmt.Println("Buah 1", buah1)
	// fmt.Println("Buah 2", buah2)
	// fmt.Println("Buah 3", buah3)

	// item1, item2, item3 := "baju", true, 1
	// fmt.Println("Item 1 : ", item1)
	// fmt.Println("Item 2 : ", item2)
	// fmt.Println("Item 3 : ", item3)

	// tipe data
	// tipe data numerik non-desimal
	println("Tipe data non-desimal")
	var nomerPositif uint8 = 89
	var nomerNegatif = -1234567890
	fmt.Printf("bilangan positif: %d \n", nomerPositif)
	fmt.Printf("Tipe data: %T \n", nomerPositif)
	fmt.Printf("bilangan Negatif: %d \n", nomerNegatif)
	fmt.Printf("Tipe data: %T \n", nomerNegatif)

	// Tipe data numerik desimal
	println("Tipe data numerik desimal")
	var phi = 3.14
	// hasil yang ditampilkan 3.140000
	fmt.Printf("Nilai phi : %f \n", phi)
	// digit desimal yang dihasilkan adalah 6 digit
	fmt.Printf("Nilai phi 2 digit : %.2f \n", phi)
	fmt.Printf("Tipe data : %T \n", phi)

	// Tipe data boolean 
	fmt.Println("Tipe data boolean")
	keluar := true
	fmt.Printf("Keluar ? %t \n", keluar)
	fmt.Printf("Tipe data : %T \n", keluar)

	// Tipe data string
	fmt.Println("Tipe data string")
	namaSaya := `nama saya Jessi Ropa,
	Umur saya 23 tahun,
	Hobi nonton youtube.`
	fmt.Println(namaSaya)
	fmt.Printf("Tipe data : %T \n", namaSaya)

	// nilai sebuah variabel ketika dideklarasi pertama kali
	var huruf string
	var angka int 
	var desimal float32
	var benar bool

	fmt.Printf("Nilai awal string %s \n", huruf)
	fmt.Printf("Tipe data: %T \n", huruf)
	fmt.Printf("Nilai awal angka %d \n", angka)
	fmt.Printf("Tipe data: %T \n", angka)
	fmt.Printf("Nilai awal desimal %f \n", desimal)
	fmt.Printf("Tipe data: %T \n", desimal)
	fmt.Printf("Nilai awal benar %t \n", benar)
	fmt.Printf("Tipe data: %T \n", benar)
}

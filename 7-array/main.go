package main

import "fmt"

func main(){
	fmt.Println("=====BELAJAR ARRAY=====")

	// var names[4] string
	// names[0] = "Jessi"
	// names[1] = "Andre"
	// names[2] = "Atri"
	// names[3] = "Reva"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// menulis array menggunakan kurung kurawal 
	// cara horizontal
	// var fruits = [4] string{"apple", "anggur", "pisang", "melon"}
	// cara vertikal 
	// var buah = [4] string{
	// 	"mangga",
	// 	"leci",
	// 	"alpucat",
	// 	"kelapa",
	// }

	// fmt.Println("Jumlah element : ", len(fruits)) // len digunakan untuk menghitung jumlah elemen sebuah array
	// fmt.Println("Isi semua element : ", fruits)
	// fmt.Println("Jumlah element : ", len(buah)) // len digunakan untuk menghitung jumlah elemen sebuah array
	// fmt.Println("Isi semua element : ", buah)

	// nilai element array tanpa jumlah element

	// var angka =[...] int {2,4,6,8,10}

	// fmt.Println("data array \t : ", angka)
	// fmt.Println("jumlah element \t : ", len(angka))

	// array multidimensi 

	// var nomor1 = [2][3]int{[3]int{3,5,7}, [3]int{2,4,6}}
	// var nomor2 = [2][3]int{{3,5,7}, {2,4,6}}
	// fmt.Println("Array 1 : ", nomor1)
	// fmt.Println("Array 1 : ", nomor2)

	// perulangan for pada array 
	// var fruits = [4] string{"apple", "anggur", "pisang", "melon"}
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Elemen ke-%d : %s\n", i, fruits[i])
	// }

	// perulangan for - range pada array
	// for i, fruit := range fruits{
	// 	fmt.Printf("Elemen ke-%d : %s\n", i, fruit)
	// }

	// Perulangan menggunakan variabel underscore pada array 
	// for _, buah := range fruits{
	// 	fmt.Printf("Nama buah : %s\n", buah)
	// }

	// alokasikan elemen array menggunakan keyword make 
	var buaah = make([]string, 2)
	buaah[0] = "sawo"
	buaah[1] = "durian"

	fmt.Println(buaah)

}
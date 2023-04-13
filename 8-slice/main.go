package main

import "fmt"

func main(){
	fmt.Println("BELAJAR SLICE")

	// inisialisasi slice
	// var hewan = []string{"kucing", "alpaca", "anjing", "kelinci"}
	// fmt.Println(hewan[0])

	// perbedaan array dan slice
	// var array1 = [2]string{"satu", "dua"}
	// var array2 = [...]int{1,2,3} 
	var slice1 = []string{"slice1", "slice2", "slice3", "slice4", "slice5"}

	// fmt.Println(array1)
	// fmt.Println(array2)
	// fmt.Println(slice1)

	// slice menggunakan operator slice 
	// var newSlice1 = slice1[0:2]
	// var newSlice2 = slice1[1:5]
	var newSlice3 = slice1[2:4]
	// var newSlice4 = slice1[0:3]
	// fmt.Println(slice1)
	// fmt.Println(newSlice1)
	// fmt.Println(newSlice2)
	// fmt.Println(newSlice3)
	// fmt.Println(newSlice4)

	// fmt.Println("\n Setelah menggantikan data \n")
	// mengganti slice 3 menjadi slice 6

	// newSlice2[1] = "slice6"
	// fmt.Println(slice1)
	// fmt.Println(newSlice1)
	// fmt.Println(newSlice2)
	// fmt.Println(newSlice3)
	// fmt.Println(newSlice4)

	// menghitung kapasitas dan panjang slice 
	fmt.Println("Slice awal : ", slice1)
	fmt.Println("Panjang slice : ", len(slice1))
	fmt.Println("Kapasitas dari slice : ", cap(slice1))

	/*
		slice1 output => [slice 1, slice 2, slice 3, slice 4, slice 5] || panjang slice : 5 || kapasitas : 5 
		newslice3 output => ---- [slice 3 ---- ---- ----] || panjang 1 || kapasitas : 4 ||
		newslice3 output => ---- ---- [slice 3 slice 4 ----] || panjang 2 || kapasitas : 3 ||
	
	*/ 
	fmt.Println("\nSlice newslice 3 : ", newSlice3)
	fmt.Println("Panjang slice : ", len(newSlice3))
	fmt.Println("Kapasitas dari slice : ", cap(newSlice3))
}
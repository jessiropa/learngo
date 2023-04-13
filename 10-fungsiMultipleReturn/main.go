package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("BELAJAR FUNGSI MULTIPLE RETURN")

	var diamater float64 = 15
	var area, keliling = kalkulasi(diamater)

	fmt.Printf("Luas lingkaran \t \t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran \t: %.2f \n", keliling)
	
}

// Membuat fungsi 
func kalkulasi(d float64) (float64, float64){
	// menghitung luas lingkaran 
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling lingkaran
	var keliling = math.Pi * d

	return area, keliling
}
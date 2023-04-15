package main

import "fmt"

func main(){
	// PERULANGAN FOR 
	// for i := 0; i <= 5; i++{
	// 	fmt.Println("Angka - ", i)
	// }

	// perulangan for dengan hanya kondisi 
	// i := 0
	// for i < 5 {
	// 	fmt.Println("angka - ", i)
	// 	i++
	// }

	// perulangan tanpa argumen 

	// for {
	// 	fmt.Println("angka - ", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// perulangan menggunakan break continue 

	// for i :=1; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("angka", i)
	// }

	// perulangan bersarang
	for i := 0; i < 5; i++{
		for j:= i; j < 5; j++{
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	
}
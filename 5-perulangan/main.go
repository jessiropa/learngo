package main

import "fmt"

func main(){
	fmt.Println("BELAJAR PERCABANGAN PADA GOLANG")

	// PERCABANGAN IF ELSE IF ELSE
	nilai := 83

	if nilai >= 90 {
		fmt.Println("score : A+")
	}else if nilai >= 85 {
		fmt.Println("score : A-")
	}else if nilai >= 75 {
		fmt.Println("score : B")
	} else if nilai >= 70 {
		fmt.Println("score : B-")
	} else if nilai >= 65 {
		fmt.Println("score : C")
	} else if nilai >= 60 {
		fmt.Println("score : C-")
	} else {
		fmt.Println("Tidak Lulus")
	}

	// VARIABEL TEMPORARY
	point := 8840.0

	if persent := point/100; persent >= 100 {
		fmt.Printf("%.1f%s Perfect \n", persent, "%")
	} else if persent >= 70 {
		fmt.Printf("%.1f%s Good \n", persent, "%")
	} else {
		fmt.Printf("%.1f%s not bad \n", persent, "%")
	}

	// SWIFT CASE
	score := 7

	switch score {
	case 8:
		fmt.Println("PERFECT")
	case 7:
		fmt.Println("AWESOME")
	default:
		fmt.Println("NOT BAD!!")
	}

	// SWIFT CASE BANYAK KONDISI
	switch score {
	case 8,9:
		fmt.Println("very perfect")
	case 5,6,7:
		fmt.Println("wow awesome")
	default:
		fmt.Println("nod bad, really")
	}

	// Kurung kurawal pada case && default 
	switch score {
	case 8,9:
		fmt.Println("very perfect")
	case 5,6,7:
		fmt.Println("wow awesome")
	default:
		{
			fmt.Println("nod bad, really")
			fmt.Println("you can be better")
		}
	}

	// swift case menggunakan gaya if else
	switch{
	case score == 8:
		fmt.Println("You perfect!!!")
	case (score <= 7) && (score >= 5) :
		fmt.Println("wow you awesome")
	default: 
		{
			fmt.Println("Not bad")
			fmt.Println("you need to learn")
		}
	}

	// swift case menggunakan fallthrough
	switch{
	case score == 8:
		fmt.Println("You perfect!!!")
	case (score <= 7) && (score >= 5) :
		fmt.Println("wow you awesome")
		fallthrough
	case score >= 4:
		fmt.Println("it's okay!!")
	default: 
		{
			fmt.Println("Not bad")
			fmt.Println("you need to learn")
		}
	}

	// percabangan bersarang 
	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!!!")
		}
	}else{
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3{
			fmt.Println("keep trying")
		}else{
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("trying harder")
			}
		}
	}
}
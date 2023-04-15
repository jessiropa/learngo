package main

import "fmt"

func main() {

	//1. Closure disimpan sebagai variabel
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numers, min, max)

	// 2. Closure sebagai nilai kembalian 
	var maks = 3
	var numbers = []int{2,3,0,4,3,2,0,4,2,0,3}
	var howMany, getNumbers = findMax(numbers, maks)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", maks)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)

}

// fungsi nilai kembalian 

func findMax(numbers []int, max int)(int, func() []int){
	var res []int
	for _, e:= range numbers{
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. Penerapan fungsi sebagai parameter

	var data = []string{"Jessi", "Kristin", "Ropa"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLenght5 = filter(data, func(each string) bool{
		return len(each) == 5
	})

	fmt.Println("data asli \t\t: ", data)
	fmt.Println("filter ada huruf \"o\" \t: ", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\" \t: ", dataLenght5)

}

// 1. penerapan fungsi sebagai parameter

// menggunakan alias skema closure
type FilterCallback func(string) bool
func filter(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
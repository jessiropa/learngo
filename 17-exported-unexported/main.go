package main

import (
	"17-exported-unexported/library"
	"fmt"
)

func main() {
	library.SayHello("jessi")
	// library.introduce("Jessi")

	var s1 = library.Student{Name: "Jessi", Grade: 23}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade ", s1.Grade)
}
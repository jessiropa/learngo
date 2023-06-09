package main

import "fmt"

func main() {
	// var s1 = student{"jessi kristin ropa", 23}
	// s1.sayHello()

	// var name = s1.getNameAt(2)
	// fmt.Println("nama panggilan :", name)

	// method pointer
	var s1 = student{"jessi kristin ropa", 23}
	fmt.Println("s1 before", s1.name)

	s1.changeName1("fujiyoshi karin")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("Harry potter")
	fmt.Println("s1 after changeName2", s1.name)
}

// 1. penerapan method

// membuat struct
type student struct {
	name  string
	grade int
}

// func (s student) sayHello() {
// 	fmt.Println("hello", s.name)
// }
// func (s student) getNameAt(i int) string{
// 	return strings.Split(s.name, " ")[i-1]
// }

// 2. method pointer

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string){
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}
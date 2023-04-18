package main

import (
	"fmt"
	"runtime"
)

// func print(till int, message string) {
// 	for i := 0; i < till; i++ {
// 		fmt.Println((i + 1), message)
// 	}
// }

// 3. channel sebagai tipe data parameter
// func printMessage(what chan string){
// 	fmt.Println(<- what)
// }
func main() {
	runtime.GOMAXPROCS(2)

	// goroutine
	// go print(5, "haiiii")
	// print(5, "gimana kabarnya ? ")
	// var input string
	// fmt.Scanln(&input)

	// 2. Channel
	// var message = make(chan string) // deklarasi variabel bertipe channel string

	// var sayHelloTo = func(who string){
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	// sedang berlangsung proses perngiriman data dari variabel kanan ke kiri
	// 	message <- data
	// }

	// go sayHelloTo("Jessi")
	// go sayHelloTo("Karin")
	// go sayHelloTo("Amyon")
	// go sayHelloTo("Taki")
	// go sayHelloTo("Hono")
	// go sayHelloTo("Ten")

	// var message1 = <- message
	// fmt.Println(message1)
	// var message2 = <- message
	// fmt.Println(message2)
	// var message3 = <- message
	// fmt.Println(message3)
	// var message4 = <- message
	// fmt.Println(message4)
	// var message5 = <- message
	// fmt.Println(message5)
	// var message6 = <- message
	// fmt.Println(message6)

	// 3. channel sebagai tipe data parameter
	// var pesan = make(chan string)

	// for _, each := range []string{"jessi", "kristin", "martha"}{
	// 	go func(who string)  {
	// 		var data = fmt.Sprintf("Hello %s", who)
	// 		pesan <- data
	// 	}(each)
	// }

	// for i := 0; i < 3; i++{
	// 	printMessage(pesan)
	// }

	// 4. eksekusi goroutine pada IIFE
	// var pesaan = make(chan string)

	// go func(who string){
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	pesaan <- data
	// }("jessi")

	// var pesann = <- pesaan
	// fmt.Println(pesann)

	// 5. Buffer channel
	pesan := make(chan int, 3)

	go func(){
		for {
			i := <- pesan
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		pesan <- i
	}
}


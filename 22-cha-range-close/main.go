package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf(" data %d", i)
	}
	close(ch)
}

func printMessage(ch <- chan string){
	for pesan := range ch {
		fmt.Println(pesan)
	}
}
func main() {
	// 1. Penerapan for range close pada chan

	runtime.GOMAXPROCS(2)

	var pesann = make(chan string)
	go sendMessage(pesann)
	printMessage(pesann)
}
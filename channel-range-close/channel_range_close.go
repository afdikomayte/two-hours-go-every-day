package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	go sendMessage(message)
	printMessage(message)

}

// ch chan string 	Parameter ch untuk mengirim dan menerima data
// ch chan<- string 	Parameter ch hanya untuk mengirim data
// ch <-chan string 	Parameter ch hanya untuk menerima datafunc
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

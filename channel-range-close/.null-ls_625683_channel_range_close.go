package main

import "fmt"

func main() {

}

func sendMessage(ch chan<- string) {

	for i := 0; i < 20; i++ {
		ch <- fmt.Sprint("data %d", i)
	}
	close(ch)
}

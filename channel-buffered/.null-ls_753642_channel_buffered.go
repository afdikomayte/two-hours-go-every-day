package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	//declared channel buffer 3
	message := make(chan int, 3)

	//gorotine func closure
	go func() {
		for {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

	//time slepp untuk membiarkan func main tetap berjalan sesuai waktu yang di tetapkan

	time.Sleep(2 * time.Second)
}

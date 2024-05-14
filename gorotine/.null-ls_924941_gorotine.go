package main

import (
	"fmt"
	"runtime"
)

func main() {
	//tentukan jumlah runtime
	runtime.GOMAXPROCS(2)

	//run function with gorotine
	go print(5, "Hello it is from print() gorotine")

	//run function usually
	print(5, "Hello there")

	var input string
	// fmt.Scanln untuk menginputkan karakter ke teminal
	// hal ini dibuat untuk menunggu proses tidak mati sebelum gorotine selesi berjalan
	fmt.Scanln(&input)

}

func print(till int, message string) {
	for i := 1; i <= till; i++ {
		fmt.Println(i, message)
	}
}

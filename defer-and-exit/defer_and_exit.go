package main

import (
	"fmt"
	"os"
)

func main() {

	//defer akan dijalankan sebelum func IIFE berakhir,
	//jika berada satu blok dengan os.exit()
	//defer tidak akan dijalankan karena program dipaksa berhenti sebelum defer dijalankan
	func() {
		// defer akan dijalankan seblum blok function selesai, dimanapun baris codenya berada
		defer fmt.Println("I am Defer")

		fmt.Println("this line was written under defer")
	}()

	os.Exit(1)
}

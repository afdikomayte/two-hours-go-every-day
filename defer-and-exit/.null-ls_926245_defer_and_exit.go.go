package main

import (
	"fmt"
	"os"
)

func main() {
	// defer akan dijalankan seblum blok function selesai, dimanapun baris codenya berada
	defer fmt.Println("I am Defer")

	fmr.Println("this line was written under defer")

	os.Exit(1)
}

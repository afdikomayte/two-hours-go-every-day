package main

import (
	"fmt"
)

func main() {

	/*
	 * constants varibles cannot be change once declared
	 */

	const name string = "variable const"

	fmt.Print("this ", name, "!\n")
	/*
	 * declaration multi const varibles
	 */
	const (
		today    string = "senin"
		sekarang        //value this variable same with today
		isToday  = true
	)

	fmt.Print(today, "\n", sekarang, "\n", isToday, "\n")
}

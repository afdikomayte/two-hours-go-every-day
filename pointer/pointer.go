package main

import (
	"fmt"
)

func main() {

	/*
	 * pointer
	 * variable pointer = var name-variable *type-data
	 * variable biasa jadi pointer = &varibale-biasa
	 */

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println(numberA, numberB)
}

func changeValueTo(number *int, toNumber int) {
	//*number -> mengambil value pointer number
	*number = toNumber
}

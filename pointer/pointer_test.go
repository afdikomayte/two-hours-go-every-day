package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer(t *testing.T) {

	//vaiable biasa
	var numberA int = 4

	//variable pointer, value pointer dari variable biasa
	var numberB *int = &numberA

	//value variable numberA sama dengan numberB
	assert.Equal(t, numberA, *numberB) //*numberB = mengambil value dari variable pointer
	fmt.Println("value numberA :", numberA)
	fmt.Println("value *numberB:", *numberB)

	//alamat memory variable numberA sama dengan variable numberB
	assert.Equal(t, &numberA, numberB) //&numberA = mengambil alamat memory(pointer) dari numberA
	fmt.Println("\naddress memory &numberA \t:", &numberA)
	fmt.Printf("address memory numberB \t\t: %v\n\n", numberB)

}

func TestChangeValuePointer(t *testing.T) {
	var numberC int = 5
	assert.Equal(t, 5, numberC)

	changeValueTo(&numberC, 15)

	assert.Equal(t, 15, numberC)
}

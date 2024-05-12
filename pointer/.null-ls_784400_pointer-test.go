package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testPointer(t testing.T) {

	//vaiable biasa
	var numberA int = 4

	//variable pointer, value pointer dari variable biasa
	var numberB *int = &numberA

	//value variable numberA sama dengan numberB
	assert.Equal(t, numberA, *numberB) //*numberB = mengambil value dari variable pointer
	//alamat memory variable numberA sama dengan variable numberB
	assert.Equal(t, &numberA, numberB) //&numberA = mengambil alamat memory(pointer) dari numberA
}

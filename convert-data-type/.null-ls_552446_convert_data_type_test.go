package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func strconv.Atoi = convert  string ke int
func TestStrconvAtoi(t *testing.T) {
	var str = "124"
	var num, err = strconv.Atoi(str)

	assert.Equal(t, nil, err)
	assert.Equal(t, 124, num)

}

// func strconv.Itoa = convert int ke string
func TestStrconvItoa(t *testing.T) {
	var num = 123
	var str = strconv.Itoa(num)

	assert.Equal(t, "123", str)

}

// func strconv.ParseInt() = convert angka string ke type numberic tertentu
// hasi dari func srrconv.ParseInt() = int64 dan error
func TestStrConvertParseInt(t *testing.T) {
	var str = "123"
	//dari string angka basis 10(desimal) lebar 64
	var num, err = strconv.ParseInt(str, 10, 64)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, int64(123), num)

	//string "1010" dikonversi ke basis 2 (biner) dengan tipe data int8 .
	strBase2 := "1010"
	numInt8, err2 := strconv.ParseInt(strBase2, 2, 8)

	assert.Nil(t, err2, "err should be nil")
	assert.Equal(t, int64(10), numInt8)
}

func TestStrconvFormatInt(t *testing.T) {
	// Berguna untuk konversi data numerik int64 ke string
	//dengan basis numerik bisa ditentukan sendiri.

	var num = int64(64)
	str := strconv.FormatInt(num, 8)

	fmt.Println(str)
}

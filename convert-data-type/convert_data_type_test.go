package main

import (
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
	str := strconv.FormatInt(num, 10)

	assert.Equal(t, "64", str)
}

// Digunakan untuk konversi float string ke float dengan lebar data bisa ditentukan.
func TestStrconvParseFloat(t *testing.T) {
	var floatStr = "24.12"
	floatEnamEmpat, err := strconv.ParseFloat(floatStr, 32)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, float64(24.1200008392334), floatEnamEmpat)
}

// func strconv.FormatFloat() mengkonfersi tipe float64 ke string dengan lebar digit desimal dan
// lebar data bisa ditentuka
func TestStrconvFormatFloat(t *testing.T) {
	numFloat64 := 12.12
	strFloat := strconv.FormatFloat(numFloat64, 'f', 4, 64)

	assert.Equal(t, "12.1200", strFloat)
}

// func strconv.ParseBool convert string bool ke bool
func TestStrconvParseBool(t *testing.T) {
	var strBool = "true"
	boolean, err := strconv.ParseBool(strBool)

	assert.Nil(t, err, "err should be nill")
	assert.Equal(t, bool(true), boolean)
}

// konversi data menggunakan teknik casting
func TestCastingIntToFloatAndRevers(t *testing.T) {
	//casting 25(int) ke tipe data float
	var numFloat float64 = float64(24)

	assert.Equal(t, float64(24), numFloat)

	//casting nilai 24.00(float32) ke int32
	var numInt32 int32 = int32(24.00)

	assert.Equal(t, int32(24), numInt32)

}

// casting string ke slice byte
func TestStringByte(t *testing.T) {
	var text = "halo"
	b := []byte(text)

	assert.Equal(t, []byte{104, 97, 108, 111}, b)

	// slice byte to string
	var bytehelo = []byte{104, 97, 108, 111}
	strhalo := string(bytehelo)

	assert.Equal(t, "halo", strhalo)
}

// mengambil nilai dari interface kosong atau any
func TestAserttion(t *testing.T) {
	var data = map[string]interface{}{
		"nama":    "afdiko",
		"age":     29,
		"height":  14.8,
		"isMale":  true,
		"hobbies": []string{"tidur", "makan"},
	}

	assert.Equal(t, "afdiko", data["nama"].(string))
	assert.Equal(t, 29, data["age"].(int))
	assert.Equal(t, 14.8, data["height"].(float64))
	assert.Equal(t, bool(true), data["isMale"].(bool))
	assert.Equal(t, []string{"tidur", "makan"}, data["hobbies"].([]string))
}

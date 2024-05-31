package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func strings.Contains => melihat apakah kata yang di parameter ke-2 ada pada parameter ke-1
func TestStringContains(t *testing.T) {
	var isExists = strings.Contains("afdiko mayte", "mayte")

	assert.Equal(t, bool(true), isExists)

}

// func strings.HasPrefix
// fungsi ini mendeteksi apakah kata pada parameter ke-1 diawali  seperti parameter ke-2
func TestStringsHasPrefix(t *testing.T) {
	var isPrefixKo = strings.HasPrefix("afdiko", "ko")

	assert.Equal(t, bool(false), isPrefixKo)

	var isPrefixAf = strings.HasPrefix("afdiko", "af")

	assert.Equal(t, bool(true), isPrefixAf)

}

// func strings.HasSuffix
// fungsi ini mendeteksi apakah akhir kata parameter ke-1 seperti parameter ke-2
func TestStringshasSuffix(t *testing.T) {
	var isSuffixKo = strings.HasSuffix("afdiko", "ko")

	assert.Equal(t, bool(true), isSuffixKo)

	var isSuffixAf = strings.HasSuffix("afdiko", "af")

	assert.Equal(t, bool(false), isSuffixAf)

}

// func strings.Count
// fungsi ini akan menghitung jumlah karakter tertentu
func TestStringsCount(t *testing.T) {
	var howMany = strings.Count("afdiko saputra", "a")

	assert.Equal(t, int(3), howMany)
}

// func strings.Index()
// untuk mencari posisi karakter pada parameter ke-2
func TestStringsIndex(t *testing.T) {
	var index = strings.Index("afdiko", "d")
	assert.Equal(t, int(2), index)

	var index2 = strings.Index("afdiko", "k")
	assert.Equal(t, int(4), index2)
}

// func strings.Replace()
// menganti string dengan parameter ke-2 dan parameter ke-3 dan parameter ke-3
// menentukan jumlah string yang akan diganti
func TestStringsReplace(t *testing.T) {
	var newText = strings.Replace("afdiko saputra", "saputra", "mayte", 1)
	assert.Equal(t, "afdiko mayte", newText)
}

// func strings.Repeat()
// parameter pertama kan di ulang sebanyak parameter ke-2
func TestStrinsRepeat(t *testing.T) {
	var strReapeat = strings.Repeat("kulu", 5)

	assert.Equal(t, "kulukulukulukulukulu", strReapeat)

}

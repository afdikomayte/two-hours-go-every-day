package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexMatchString(t *testing.T) {
	var text = "afdiko mayte saputra"
	// membuat aturan regex nya
	var regex, err = regexp.Compile(`[a-z]+`)

	// melihat apakah kalimat cocok dengan aturan regex yang di buat
	var isMatch = regex.MatchString(text)

	assert.Nil(t, err, "Should be nil")

	assert.Equal(t, bool(true), isMatch)
}

//fungsi lain nya pun sama cara penggunaannya

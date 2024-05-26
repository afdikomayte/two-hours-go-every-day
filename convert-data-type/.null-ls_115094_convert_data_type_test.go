package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrconvAtoi(t *testing.T) {
	var str = "124"
	var num, err = strconv.Atoi(str)

	assert.Equal(t, nil, err)
	assert.Equal(t, 124, num)

}

func TestStrconvItoa(t *testing.T) {
	var num = 123
	var str = strconv.Itoa(num)

	assert.Equal(t, "123", str)

}

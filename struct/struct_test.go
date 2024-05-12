package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruct(t *testing.T){
  //make object (object adalah variale dari struct)
  var employe1 Person
  employe1.name = "Afdiko"
  employe1.gender = "Male"
  employe1.age = 29

  assert.Equal(t, "Afdiko",employe1.name)
  assert.Equal(t, "Male",employe1.gender)
  assert.Equal(t, uint(29), employe1.age)
}

func TestGetLencana(t *testing.T){
  
  employe2 := Person{
    name: "Afdiko",
    gender: "Male",
    lencana: "Belum Ada",
    absent: 21,
    age: uint(29),
  }

  GetLencana(&employe2)

  fmt.Printf("\n %v \n\n",employe2)

  assert.Equal(t, "Teladan", employe2.lencana)
}

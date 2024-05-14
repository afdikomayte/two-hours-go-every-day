package main

import (
	"testing"
  "fmt"

	"github.com/stretchr/testify/assert"
)

func TestPersegi(t *testing.T){

  //bangun datar persegi 
  var bangunDatar hitung
  
  bangunDatar = persegi{10.0}
  assert.Equal(t,100.0,bangunDatar.luas())
  assert.Equal(t, 40.0,bangunDatar.keliling())

}

func TestLingkaran(t *testing.T){

  var bangunDatarLingkaran hitung 
 
  bangunDatarLingkaran = lingkaran{14}

  fmt.Println(bangunDatarLingkaran.luas())
  fmt.Println(bangunDatarLingkaran.keliling())
  fmt.Println(bangunDatarLingkaran.(lingkaran).jariJari())
}

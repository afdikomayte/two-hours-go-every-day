package main

import (
	"math"
)

func main(){
}

//interface hitung
type hitung interface{
  luas() float64
  keliling() float64
}

//struct lingkaran
type lingkaran struct{
  diameter float64
}

func (l lingkaran) jariJari() float64{
  return l.diameter / 2
}

func (l lingkaran) luas() float64{
  return math.Pi*math.Pow(l.jariJari(),2)
}

func (l lingkaran) keliling() float64{
  return math.Pi * l.diameter
}

//struct persegi
type persegi struct{
  sisi float64
}

func (p persegi) luas() float64{
  return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64{
  return p.sisi * 4
}





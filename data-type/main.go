package main

import "fmt"

func main(){

  /*
  * Data type numeric non floating point
  *
  * uint -> positive
  * int -> negative and positive
  *
  * byte -> same uint8
  * rune -> same int32
  */
  
  var positiveNumber uint8 = 89
  var negativeNumber int32 = -1243423644

  fmt.Printf("bilangan positive: %d\n", positiveNumber)
  fmt.Printf("bilangan negative: %d\n", negativeNumber)

  /*
  * Data type numeric floting point
  * float32 and float64
  */

  decimalNumber := 2.64

  fmt.Printf("bilangan desimal=%f\n",decimalNumber)
  fmt.Printf("bilangan desimal=%.3f\n",decimalNumber)

  /*
  * Data type -> bool (Boolean)
  */

  var exist bool = true
  fmt.Printf("exist ? % t \n", exist)

  /*
  * Data type -> string
  */

  var intro string = `Perkenalkan Nama saya "Afdiko".
  Mari Kira belajar Golang`
  fmt.Println(intro)

  /*
  * Data type can be nil
  * pointer
  * func 
  * slice
  * map[type]type
  * channel
  * interface kosong atau any (alias dari -> interface{})
  */

}

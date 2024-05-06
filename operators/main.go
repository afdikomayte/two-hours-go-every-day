package main

import "fmt"

func main(){

  /*
  * Arithmetic Operators
  *
  * +,-,*,/,%
  */ 
  
  var value = (((2+6)%3)*4-2)/3

  fmt.Println("(((2+6)%3)*-2)/3 =",value)

  /*
  * Operator Perbandingan
  * == sama dengan
  * != tidak sama dengan
  * < lebih kecil
  * <=lebih kecil atau sama dengan
  * > lebih besar
  * >= lebih besar atau sama dengan
  */
  
  var isEqual = (value == 2)

  fmt.Print(isEqual)

  /*
  * Operator Logika
  * && dan 
  * || atau 
  * ! kebalikan/negasi
  */

  var left = false
  var right = true

  var leftAndRight = left && right
  fmt.Printf("left && right \t(%t) \n", leftAndRight)

  var leftOrRight = left || right
  fmt.Printf("left || right \t(%t) \n", leftOrRight)

  var leftReverse = !left 
  fmt.Printf("!left \t(%t) \n", leftReverse)

}

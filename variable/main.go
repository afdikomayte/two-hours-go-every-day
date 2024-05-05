package main

import "fmt"

func main(){

  /*
  *make variable with var schema
  *formula
  *
  * var <name-variable> <tipe-data> = <value-variable>
  * 
  * if you want make variable without value, just take off value 
  * and you can insert value later
  */

  var firstName string = "Afdiko"

  var lastName string
  lastName = "Mayte"

  //use fmt.Printf for print with format
  fmt.Printf("My Name is %s %s\n", firstName, lastName)

}

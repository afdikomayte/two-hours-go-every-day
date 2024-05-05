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

  /*
  * variable without data-type
  * formula 
  * <variable-name> := <value>
  */

  fullName := "Afdiko Mayte Saputra"
  fmt.Printf("My fullname =%s\n",fullName)

  /*
  * declaration multi-variable
  */ 
  var fruit1,fruit2, fruit3 = "manggo","apple","strobbery"
  // or
  animal1, animal2, animal3 := "dog","cat","tiger"
  
  fmt.Println(fruit1, fruit2, fruit3)
  fmt.Println(animal1,animal2,animal3)


  /*
  * variable _ (underscore) in go all variable was declaration must be use
  * if variable not be use make variable name with _ (underscore)
  */
  useVaribale, _ := "use-variable","ignore-variable"

  fmt.Println(useVaribale)

  /*
  * declaration variable function new()
  * variables declared with function new() become pointer variable 
  * formula
  * name-variable := new(data-type)
  */
  
  varPointer := new(string)

  fmt.Printf("It's Variable-pointer = %x\n", varPointer)
  fmt.Printf("It's Value Varible ponter = %s\n",*varPointer)

}

package main

import "fmt"

func main(){

  /*
  * array
  */ 

  var names [3]string
  names[0] = "afdiko"
  names[1] = "mayte"
  names[2] = "saputra"

  fmt.Println(names[0],names[1],names[2])

  var fruits = [3]string{"apple","grape","banana"}

  fmt.Println("Jumlah element \t\t", len(fruits))
  fmt.Println("Isi semua element \t", fruits)
  
  //inisialisasi tanpa Jumlah element
  numbers := [...]int{1,2,3,4,5}
  fmt.Println("Angka  \t\t\t",numbers)

  fmt.Println("\nArray Multi Dimensi")
  numbers1 := [2][3]int{
    {1,2,3},
    {4,5,6},
  }
  fmt.Println("numbers1 \t\t",numbers1)

  /*
  * loops array with for
  */
    
  fmt.Println("\nLoops array with for")

  animals := [3]string{"dog","tiger","snack"}

  for i := 0; i < len(animals); i++{
    fmt.Printf("index %d: \t %s\n",i,animals[i])
  }

  /*
  * loops array with for-range
  */
  fmt.Println("\nLoops array with for-range")
  for i,v := range animals{
    fmt.Printf("Binatang %d: \t %s\n",(i+1),v)
  }

  /*
  * declaration array with make() keyword
  * var variable-name = make([]data-type, len-array)
  */

  fmt.Println("\nDeclaratio array with keyword make()")
  var animals2 = make([]string, 3)
  animals2[0] = "fish"
  animals2[1] = "bison"
  animals2[2] = "butterflay"

  fmt.Println("Binatang : \t",animals2)

}

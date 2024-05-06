package main

import "fmt"

func main(){
  
  /*
  * for 
  */
  for i := 0; i < 5; i++{
    fmt.Println("Ke-",i)
  }

  /*
  * for with just conditional argument
  */

  mulai := 1
  
  fmt.Println("\nMulai dari 1")
  for mulai <= 5 {
    fmt.Println("Angka ", mulai)
    mulai++
  }

  /*
  * for without conditional argument, end loop with if and break
  */

  angka := 1

  fmt.Println("\nAngka 1 - 5")

  for{
    fmt.Println("Angka ", angka)

    angka++
    if angka == 6 {
      break
    }
  }

  /*
  * for-range
  */

  fmt.Println("\n for range")

  var xs string= "STRING"

  fmt.Println("letters of STRING")
  for in, va := range xs{
    fmt.Println("index :",in," letters:",string(va))
  }
  
  /*
  * for with break and continue
  */

  fmt.Println("\n for with break and continue")

  for number := 1; number <= 10; number++{
    if number % 2 == 0{
      continue
    }

    if number > 8 {
      break
    }

    fmt.Println("Angka ganjil < 8 ke :",number)
  }

  /*
  * nested loops
  */

  fmt.Println("\nPerulangan bersarang")
  

  for i2 := 4; i2 >= 1; i2--{
    for j2 := i2; j2 <= 4; j2++{
      fmt.Print("*"," ")
    }
    fmt.Println()
 }
  for i1 := 1; i1 <= 5; i1++{
    for j1 := i1; j1 <= 5; j1++{
      fmt.Print("*"," ")
    }
    fmt.Println()
  }
  
  /*
  * for with label
  */
  fmt.Println("\nFor with label to break")
  outerLoop:
  for i := 1; i < 5; i++{
    for j := 0; j < 5; j++{
      if i == 3{
        break outerLoop
      }
      fmt.Print("matriks[",i,"][",j,"]","\n")
    }
  }
}

package main

import "fmt"

func main(){
  
  /* conditional statement
  *if 
    * if condition {
    *     
    * } else if condition{
    *
    * }
  * 
  *switch 
    *switch expression {
    *case condition:
    *   
    *}
  */ 
  
  var point = 10

  if point == 10{
    fmt.Println("Lulus dengan sempurna")
  }else if point >=5 {
    fmt.Println("lulus")
  }else{
    fmt.Println("gagal")
  }

  var level = 5
  
  switch level{
  case 5:
    fmt.Println("Level 5 -> anda diizinkan masuk")
  case 1,2,3,4:
    fmt.Println("Level Anda belum memenuhi syarat")
  default:
    fmt.Println("Anda bukan pemain")
  } 
  

}

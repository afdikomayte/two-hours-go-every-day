package main

import "fmt"

func main(){

  /*
  * map 
  */

  //insiasi 
  var fruits = map[string]string{"pinnaple":"7 kg"}
  
  //add key = value
  fruits["apple"] = "3 kg"
  fruits["grape"] = "5 kg"

  fmt.Printf("jumlah pinnaple: %s\t\n",fruits["pinnaple"])
  fmt.Printf("jumlah apple : %s\t\n",fruits["apple"])
  fmt.Printf("jumlah grape : %s\t\n",fruits["grape"])

  //insiasi with make 
  var animals = make(map[string]string) 
  animals["dog"] = "puppy"
  animals["cat"] = "catty"

  fmt.Printf("Peliharaan 1 dog : nama %s\n",animals["dog"])
  fmt.Printf("Peliharaan 2 cat : nama %s\n",animals["cat"])
  
  /*
  * iterasi map with for-range
  */
  fmt.Println("\nIterasi map dengan for-range")
  for key, value := range fruits{
    fmt.Print(key,":",value,"\n")
  }

  /*
  * delete element 
  */

  fmt.Println("\nDelete element map")
  
  delete(fruits,"pinnaple");

  fmt.Println(fruits)

  /*
  * check value map 
  */

  fmt.Println("\nChecking map")
  value, isExist := fruits["pinnaple"]

  if isExist {
    fmt.Println(value)
  }else{
    fmt.Println("Item is not exists")
  }

  /*
  * combination slice & map 
  */

  fmt.Println("Kombinasi slice dan map (map dalam slice)")

  var students = []map[string]string{
    {"name":"afdiko","gender":"male"  },
    {"name":"anisa", "gender":"female"},
  }

  for _, student := range students{
    for key, value := range student{
      fmt.Println(key ,"\t:",value)
    }
    fmt.Println()
  }

}

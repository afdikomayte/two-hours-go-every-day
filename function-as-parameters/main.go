package main 

import (
  "fmt"
  "strings"
)

func main(){
  var datas = []string{
    "kamu",
    "kami",
    "dia",
  }
  //closure checkWordContainsI
  checkWordContainsI := func(word string) bool{
    return strings.Contains(word,"i")
  } 

  //kata yang mengandung i 
  var wordContainsI = filter(datas, checkWordContainsI)
  
  fmt.Printf("datas \t\t: %v \n", datas)
  fmt.Printf("word contains i : %v \n",wordContainsI)

}

//functian filter strings 
func filter(datas []string, callback func(string) bool) []string{
  var result []string

  for _, each := range datas{

    if filtered := callback(each); filtered{
      result = append(result,each)
    }
  }

  return result
}

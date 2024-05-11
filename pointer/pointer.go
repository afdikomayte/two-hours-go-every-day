package main

import "fmt"


func main(){
  var name = "afdiko"
  sayHello(name)
}

func sayHello(name string)  {
  fmt.Println("Hello ",name)
}

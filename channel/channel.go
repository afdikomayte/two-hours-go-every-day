package main

import (
	"fmt"
	"runtime"
)


func main(){
  runtime.GOMAXPROCS(2)

  var message = make(chan string)

  var sayHello = func(name string){
    var data = fmt.Sprintf("Hello %s", name)

    message <- data
  }

  go sayHello("Afdiko")
  go sayHello("Mayte")

  var message1 = <- message
  fmt.Println(message1)

  var message2 = <- message
  fmt.Println(message2)
   

  fmt.Println("\n channel as arguments")
  //channel as arguments
  var messageChan = make(chan string)

  //perulangan gorotine closure fungsi say package 
  for _, each := range []string{"afdiko", "mayte", "saputra"}{
    go func(name string){
      var hi = fmt.Sprintf("Hi %s", name)
      messageChan <- hi
    }(each)
  }

  //Println channel melalui function print channel
  for i := 1; i <= 3; i++{
    printChan(messageChan)
  }

}

// fungsi untuk print channel
func printChan(messageChannel chan string){
  fmt.Println(<-messageChannel)
}



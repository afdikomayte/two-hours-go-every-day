package main

import "fmt"


func main(){

  numbers := []int{2,3,6,4,7,5,2,4,5}
  fmt.Printf("numbers\t :%v \n", numbers)

  //make channel ch1 for getAverage 
  var ch1 = make(chan float64)
  go getAverage(numbers,ch1)

  //make channel ch2 for getMax
  var ch2 = make(chan float64)
  go getMax(numbers, ch2)

  for i = 1; i <= 2; i++{
    switch{
    case avg := <- ch1:
      fmt.Printf("average\t :%v \n", avg)
    case max := <- ch2:
      fmt.Prprint("max\t :%v \n",max)
    }
  }

 }

func getAverage(numbers []int, ch chan float64){
  
  var sum = 0

  for _, number := range numbers{
    sum += number
  }

  ch <- float64(sum)/float64(len(numbers))
}

func getMax(numbers []int, ch chan int){

  var max = 0

  for _, v := range numbers{
    if v <= max {
      continue
    }
    max = v
  }

  ch <- max
}

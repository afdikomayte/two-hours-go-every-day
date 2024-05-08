package main

import "fmt"

func main(){
  student := "Afdiko"

  avg, greet:=calculateAndGreeting(student, 2,3,4,3)

  avgToString := fmt.Sprintf("Rata-rata : %.2f",avg)
  fmt.Println(avgToString)
  fmt.Println(greet)

  // pengisian variable varadic dengan slice
  fmt.Println("\nPengisian argument varadic dengan variable slice")

  nilai := []int{2,3,6,7,3}
  avg2, _ := calculateAndGreeting(student,nilai...)
  avg2ToString := fmt.Sprintf("Rata-rata : %.2f",avg2)
  fmt.Println(avg2ToString)

}

//argument numbers -> variable biasa, students -> slice
func calculateAndGreeting(student string, numbers ...int)(float64, string){
  //calculate 
  var total int = 0

  for _,number := range numbers{
    total += number
  }

  var avg = float64(total)/float64(len(numbers)) //float64(int) -> casting int to float64
  
  //greeting 
  
  greet := "Halo, " + student
  

  return avg, greet
}


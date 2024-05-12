package main


func main(){
 /*
  * struct
  * type struct-name struct{
    *
    *
  *}
  */  

  
}

 //inisiasi struct
type Person struct{
  name string
  gender string
  lencana string
  age uint
  absent int
}

func GetLencana(employe *Person) *Person{

  if employe.absent >= 20 {
    employe.lencana = "Teladan"
  }else if employe.absent >= 15 {
    employe.lencana = "Rajin"
  }else{
    employe.lencana = "Biasa"
  }

  return employe

}

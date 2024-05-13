package method

import (

	"strings"
)

func main(){

 
}

var Gender = []string{"Laki-laki","Perempuan"}

type Student struct{
  name string
  gender string
  class int
}

func(s *Student) SetName(name string, gender int, class int){
  s.name = name
  s.gender = Gender[gender]
  s.class = class
}


func (s Student) GetName() string{
  return strings.Split(s.name, " ")[0]
}




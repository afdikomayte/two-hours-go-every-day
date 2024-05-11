package main

import "testing"
import "reflect"
//import "fmt"

func TestRandomWithRange(t *testing.T){
  result := randomWithRange(2,9)
  reflectResult := reflect.ValueOf(result)
  
  if reflectResult.Kind() != reflect.Int {
    panic("result type must int")
  }
  
}

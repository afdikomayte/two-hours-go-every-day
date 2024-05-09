package main

import "fmt"

func main(){

  var getMinMax = func(n []int) (int,int){
    var min, max int

    for i, v := range n{
      switch {
        case i == 0:
          min, max = v,v 
        case v > max:
          max = v 
        case v < min:
          min = v 
        }
    }
    return min, max
  }

  // make slice as argument
  numbers := []int{2,2,3,4,5,8,4,3,1}
  min, max := getMinMax(numbers)

  //%v for display any data-type
  fmt.Printf("numbers :%v \nmin : %v \nmax: %v\n",numbers,min,max)

  /*
  * Immediately-Invoked Function Expression (IIFE)
  *
  * Closure jenis IIFE ini eksekusinya adalah langsung saat deklarasi
  */

  fmt.Printf("\nClosure jenis IIFE ini eksekusinya adalah langsung saat deklarasi")

  fmt.Println("\nFilter numbers great then 3")

  numbers2 := []int{2,2,3,4,6,0,1,2,3,4}

  filteredNumbers := func(minNumber int) []int{
    var r []int
    for _,e := range numbers2{
      if e < minNumber{
        continue
      }
      r = append(r,e)
    }

    return r 
  }(3) //untuk langsung meng-eksekusi
  
  fmt.Printf("numbers : \t%v \nfiltered numbers (min3) : %v \n",numbers2,filteredNumbers)

  /*
  * closure as return value
  */
  fmt.Println("\nfind numbers less then 4")
  numbers5 := []int{2,3,4,5,6,7,8,4}

  totalNumbers, getMaxNumbers := findMax(numbers5,4)
  numbersLess4 := getMaxNumbers()

  fmt.Println("numbers\t:", numbers5)
  fmt.Printf("find \t: 4\n\n",)

  fmt.Println("jumlah number di bawah 4:", totalNumbers)
  fmt.Println("number di bawah 4 \t:", numbersLess4)


}


/*
* closure as return value
*/

func findMax(nm []int, max int) (int, func()[]int){
  var numbers4 []int

  for _, v := range nm{
    if v <= max{
      numbers4 = append(numbers4,v)
    }
  }

  return len(numbers4), func ()[]int  {
    return numbers4
  }


}

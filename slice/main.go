package main

import "fmt"

func main(){

  /*
  * slice
  * slice doesn't has len when declation
  */
  fmt.Println("Slice")
  
  var fruits = []string{"apple","banana","grape"}
  fmt.Println(fruits)

  /*
  * make slice from array
  */
  fmt.Println("\nSlice from array")
  var arrayFruits = [3]string{"apple","banana","potato"}

  var sliceFromArrayFruits1 = arrayFruits[0:2] //make slice from array index 0 to < 2 (0,1)
  
  fmt.Println("array fruits : \t\t\t",arrayFruits)
  fmt.Println("len arrayFruits: \t\t", len(arrayFruits))
  fmt.Println("capacity arrayFruits: \t\t",cap(arrayFruits))

  fmt.Println("slice from arrayFruits:\t\t",sliceFromArrayFruits1)
  fmt.Println("len sliceFromArrayFruits1:\t",len(sliceFromArrayFruits1))
  fmt.Println("capacity sliceFromArrayFruits1:\t", cap(sliceFromArrayFruits1))

  /*
  * Slice is data type reference
  * so if you make change in slice from array or other slice
  * all data with same reference will also change
  */
  fmt.Println("\nMerubah isi slice dari array atau slice lainya akan merubah data referensinya")

  var sliceFromArrayFruits2 = arrayFruits[0:2]
  var sliceFromArrayFruits3 = arrayFruits[1:2]

  fmt.Println(arrayFruits)
  fmt.Println(sliceFromArrayFruits2)
  fmt.Println(sliceFromArrayFruits3) 
  
  fmt.Println(`"sliceFromArrayFruits3[0] = "pinnaple"`)

  sliceFromArrayFruits3[0] = "pinnaple"
  fmt.Println(arrayFruits)
  fmt.Println(sliceFromArrayFruits2)
  fmt.Println(sliceFromArrayFruits3) 

  /*
  * fuction append()
  * add new element to end of slice
  */ 

  fmt.Println("\nfungsi append untuk menambah element slice setelah index terakhir")
  fmt.Println("arrayFruits:",arrayFruits)
  fmt.Println("slice sumber sliceFromArrayFruits4: ",sliceFromArrayFruits1)
 
  fmt.Println("\nsetelah dilakukan append() pada sliceFromArrayFruits4(ini slice baru)")

  sliceFromArrayFruits4 := append(sliceFromArrayFruits1,"tomato")

  fmt.Println("arrayFruits:",arrayFruits)
  /*
  * potato pada index terakhir pada array akan beruah menjadi tomato
  * karena panjang sliceFromArrayFruits1 reference sliceFromArrayFruits4 hanya 2 
  * sedang kan index arrayFruits sumber awal slice 3 
  * sehingga di lakukan append pada sliceFromArrayFruits4 
  * sliceFromArrayFruits4 akan menambah index 3 
  * dan pada arrayFruits index 3 akan berubah 
  */
  fmt.Println("slice sumber sliceFromArrayFruits4: ",sliceFromArrayFruits1)
  fmt.Println("sliceFromArrayFruits4:",sliceFromArrayFruits4)
  

 /*
 * function copy()
 */
 fmt.Println("\nfungsi copy")
 sliceDst := make([]string, 2)
 sliceSrc := arrayFruits[0:3]
 copy(sliceDst, sliceSrc)

 fmt.Println("slice sumber len 3 :", sliceSrc)
 fmt.Println("slice tujuan len 2 :", sliceDst)
 fmt.Println("element yang tercopy hanya sesuai panjang slice tujuan")






}

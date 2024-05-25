package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main(){
  
  //user akan diminta memasukan nama
  //jika user langsung menekan enter program akan error
  //jika user menekan masukan angka program akan panic
  //program yang panic akan di recover dengan func catch yang kita buat
  
  var input string

  fmt.Print("Masukan Nama Anda\t:")
  // membaca input terminal bisa memasukan lebih dari 1 kata yang di pisahkan space
  reader := bufio.NewReader(os.Stdin)
  input,_ = reader.ReadString('\n')
  
  //hapus space di awal dan akhir
  trimmedInput := strings.TrimSpace(input)
 
  // validate inputan sesuai ketentuan yang telah di tentukan func validate
  if valid, err := validate(trimmedInput); valid{
    fmt.Println("halo", trimmedInput)
  }else{
    fmt.Println(err.Error())
  }
}

func catch(){

  if r := recover(); r != nil{
    fmt.Println("terjadi error")
  }
}

//func validate akan memvalidasi inputan user jika white space saja atau ada ngaka program akan error
func validate(input string)(bool, error){
  //splite input untuk cek angka
  sliceInput := strings.Split(input,"")
  
  numberFound := false

  //lakukan pengecekan input kurang dari 3 karakter
  if len(sliceInput) < 3{ 
    // lakukan perulangan pada sliceInput, jika ada angka, rubah numberFound menjadi true
    loop:
    for _, e := range sliceInput{
      for _, i := range []string{"0","1","2","3","4","5","6","7","8","9"}{
        if e == i{
          numberFound = true
          break loop
        }
      } 
    }
  }
  if len(sliceInput) < 3{
    return false, errors.New("Nama tidak boleh kosong atau kurang 3 karakter")
  }else if numberFound == true {
    return false, errors.New("Nama tidak boleh mengandung angka")
  }
  return true, nil
}

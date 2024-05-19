package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	//buat channl untuk media kirim dan terima data antara gorotine
	var chanData = make(chan map[string]string)

	go sendData(chanData)
	retreiveData(chanData)

}

func sendData(ch chan<- map[string]string) {
	//buat randomizer untuk generate waktu acak
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	var datas = make(map[string]string)

	for i := 1; true; i++ {
		currentTime := time.Now()
		formattedTime := currentTime.Format("15:04:05")

		datas["data"] = strconv.Itoa(i)
		datas["time"] = formattedTime

		ch <- datas

		//for akan berhenti sleep mengirim data dengan waktu acak
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan map[string]string) {

loop:
	for {
		select {
		case data := <-ch:
			fmt.Printf("Receive data : %v \t\t\t-> time : %v \n", data["data"], data["time"])
		case <-time.After(time.Second * 5):

			currentTime := time.Now()
			formattedTime := currentTime.Format("15:04:05")

			fmt.Printf("timeout. no activities under 5 seconds \t-> time : %v \n", formattedTime)
			break loop
		}
	}

}

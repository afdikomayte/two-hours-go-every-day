package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Membuat scanner untuk membaca input dari terminal
	var scanner = bufio.NewScanner(os.Stdin)
	// Mencetak pesan permintaan input
	fmt.Print("masukan angka (pisahkan dengan space) \t:")
	// Memanggil scanner untuk membaca input
	scanner.Scan()
	//mengambil inputan
	input := scanner.Text()
	//remove whiteSpace in the end input
	input = strings.TrimSpace(input)
	// split intputan yang di pisahkan space dan menghasilkan []string
	numbersString := strings.Split(input, " ")

	//conver []string ke []int
	var numbers = []int{}
	for _, numStr := range numbersString {
		//conversi angka string ke int
		num, _ := strconv.Atoi(numStr)

		//strconv.Atoi() akan menghasilkan 0 jika selain string angka
		//skip 0 tersebut
		if num == 0 {
			continue
		}

		//masukan
		numbers = append(numbers, num)
	}

	//numbers := []int{2, 3, 6, 4, 7, 5, 2, 4, 5}
	fmt.Printf("numbers\t\t\t\t\t:%v \n", numbers)

	//make channel ch1 for getAverage
	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	//make channel ch2 for getMax
	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 1; i <= 2; i++ {
		//use select for select channel from gorotine
		select {
		case avg := <-ch1:
			fmt.Printf("average\t\t\t\t\t:%v \n", avg)
		case max := <-ch2:
			fmt.Printf("max\t\t\t\t\t:%v \n", max)
		}
	}

}

func getAverage(numbers []int, ch chan float64) {

	var sum = 0

	for _, number := range numbers {
		sum += number
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {

	var max = 0

	for _, v := range numbers {
		if v <= max {
			continue
		}
		max = v
	}

	ch <- max
}

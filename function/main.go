package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("func return value ")
	randomValue := randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)

	randomValue = randomWithRange(3, 8)
	fmt.Println("random number :", randomValue)

	fmt.Println("\nreturn to stop proses")
	dividenNumber(10, 2)
	dividenNumber(5, 0)
	dividenNumber(100, 3)

}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

/*
* function and return value
 */
func randomWithRange(min, max int) int {

	var value = randomizer.Int()%(max-min+1) + min

	return value

}

/*
* return to stop proses in block code
 */

func dividenNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return //to stop proses cause this func is void
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}

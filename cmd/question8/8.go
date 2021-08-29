package main

import (
	"fmt"
	"math"
)

/*
Дана переменная int64. Написать программу которая устанавливает i-й бит в 1
или 0
*/

func main() {
	var (
		n    int64 = 1000
		pos  int64 = 2
		mask int64 = 1 << pos
	)

	fmt.Printf("%064b %d\n", n, n)
	if n >= 0 {
		n = n ^ mask
	} else {
		n--
		n = math.MaxInt64 &^ n
		n = (n ^ mask) * -1
	}

	fmt.Printf("%064b %d\n", n, n)
}

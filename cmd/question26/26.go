package main

import (
	"fmt"
)

/*
Написать программу, которая переворачивает строку. Символы могут быть
unicode
*/

func main() {
	var (
		str string = "dlrow olleh"
	)
	rs := []rune(str)

	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-i-1] = rs[len(rs)-i-1], rs[i]
	}

	fmt.Println(string(rs))
}

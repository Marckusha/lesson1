package main

import (
	"fmt"
	"strings"
)

/*
Написать программу, которая переворачивает слова в строке (snow dog sun -
sun dog snow)
*/

func main() {
	var (
		str    = "snow dog sun"
		result string
	)

	split := strings.Split(str, " ")

	for i := 0; i < len(split)/2; i++ {
		split[i], split[len(split)-i-1] = split[len(split)-i-1], split[i]
	}

	result = strings.Join(split, " ")

	fmt.Println(result)
}

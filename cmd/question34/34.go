package main

import (
	"fmt"
	"strings"
)

/*
Написать программу, которая проверяет, что все символы в строке уникальные
*/

func uniqueStr(str string) bool {
	for _, i := range str {
		if strings.Count(str, string(i)) > 1 {
			return false
		}
	}

	return true
}

func main() {
	str := "murka"
	fmt.Println(uniqueStr(str))
}

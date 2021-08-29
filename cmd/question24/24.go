package main

import (
	"fmt"
)

/*
Создать слайс с предварительно выделенными 100 элементами
*/

func main() {
	slice := make([]int, 100, 100)
	fmt.Println(len(slice))
}

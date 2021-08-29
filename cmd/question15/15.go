package main

import (
	"fmt"
)

/*
Поменять местами два числа без создания временной переменной
*/

func main() {
	var (
		x int = 30
		y int = -99
	)

	//old version
	x += y
	y = x - y
	x = x - y

	//new version
	//x, y = y, x
	fmt.Println(x, y)
}

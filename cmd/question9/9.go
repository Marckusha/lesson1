package main

import (
	"fmt"
)

/*
Написать конвейер чисел. Даны 2 канала - в первый пишутся числа из массива,
во второй пишется результат операции 2*x, после чего данные выводятся в
stdout
*/

func sendingFunc(sendChan chan int, array []int) {

	for _, elem := range array {
		sendChan <- elem
	}
	close(sendChan)
}

func calculatingFunc(sendChan, calcChan chan int, x int) {
	for elem := range sendChan {
		calcChan <- elem * x
	}
	close(calcChan)
}

func main() {

	var (
		array     = []int{1, 2, 3, 4, 5, 6}
		x     int = 99
	)

	sendChan := make(chan int)
	calcChan := make(chan int)

	go sendingFunc(sendChan, array)
	go calculatingFunc(sendChan, calcChan, x)

	for c := range calcChan {
		fmt.Println(c)
	}
}

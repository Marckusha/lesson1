package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Даны 2 канала - в первый пишутся рандомные числа после чего они
проверяются на четность и отправляются во второй канал. Результаты работы
из второго канала пишутся в stdout
*/

func generator(out chan<- int) {
	rand.Seed(time.Now().UnixNano())

	for {
		out <- rand.Intn(100)
		time.Sleep(time.Millisecond * 200)
	}
}

func filter(in <-chan int, out chan<- int) {
	for {
		if val, _ := <-in; val%2 == 0 {
			out <- val
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go generator(ch1)
	go filter(ch1, ch2)

	for {
		select {
		case v := <-ch2:
			fmt.Println(v)
		}
	}

}

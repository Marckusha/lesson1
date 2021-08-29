package main

import (
	"context"
	"fmt"
	"time"
)

/*
Написать программу, которая будет последовательно писать значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна
завершиться.
*/

func Write(ctx context.Context, channel chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			{
				channel <- i
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
}

func Read(ctx context.Context, channel <-chan int) {
	for {
		select {
		case x := <-channel:
			fmt.Println(x)
		case <-ctx.Done():
			return
		}
	}
}

func main() {

	var N time.Duration
	fmt.Scan(&N)

	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	ch := make(chan int)

	go Write(ctx, ch)
	Read(ctx, ch)

	fmt.Print("end")
}

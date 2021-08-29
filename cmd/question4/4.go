package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

/*
Реализовать набор из N воркеров, которые читают из канала произвольные
данные и выводят в stdout. Данные в канал пишутся из главного потока.
Необходима возможность выбора кол-во воркеров при старте, а также способ
завершения работы всех воркеров
*/

func funcWorker(ctx context.Context, channel <-chan string, idWorker int) {
	for {
		select {
		case v := <-channel:
			fmt.Println("id worker:", idWorker, v)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	var N int
	fmt.Print("Enter count workers: ")
	fmt.Scan(&N)

	shutDown := make(chan struct{}, 1)

	channel := make(chan string, N)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < N; i++ {
		go funcWorker(ctx, channel, i)
	}

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			str := scanner.Text()
			if str == "exit" {
				shutDown <- struct{}{}
				return
			}
			channel <- str
		}
	}()

	<-shutDown
	cancel()
	fmt.Println("end")
}

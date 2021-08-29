/*
1. при отмене контекста exersize 4.go
2. при закрытии канала 6, 6_2
3. создать небуф канал, который будет являться закрывающим 6_3
*/

package main

import (
	"fmt"
	"time"
)

/*
Какие существуют способы остановить выполнения горутины? Написать
примеры использования
*/

func main() {
	stopCh := make(chan int)

	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				fmt.Println("todo smth")
			}
		}
	}()

	time.Sleep(time.Second * 1)
	close(stopCh)
}

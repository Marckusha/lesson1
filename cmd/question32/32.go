package main

import (
	"fmt"
	"time"
)

func Sleep(t time.Duration) {
	select {
	case <-time.After(t):
		return
	}
}

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	Sleep(time.Millisecond * 500)
	fmt.Println("end")
}

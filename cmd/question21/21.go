package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Написать программу, которая в конкурентном виде читает элементы из массива
в stdout.
*/

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		wg          sync.WaitGroup
		min             = 10
		max             = 20
		sizeSlice       = rand.Intn(max-min) + min
		arr             = []int{}
		countReader int = 4
	)

	for i := 0; i < sizeSlice; i++ {
		arr = append(arr, rand.Intn(100))
	}

	for i := 0; i < countReader; i++ {
		wg.Add(1)
		go func(num int) {
			for i := 0; i < len(arr); i++ {
				fmt.Printf("num: %d elem: %d\n", num, arr[i])
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

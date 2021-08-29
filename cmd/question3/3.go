package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел (2,4,6,8,10) найти их сумму
квадратов(2
2+3
2+4
2….) с использованием конкурентных вычислений.
*/

func main() {
	var (
		array  = []int32{2, 4, 6, 8, 10}
		wg     sync.WaitGroup
		result int32
	)

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(num int32) {
			defer wg.Done()
			atomic.AddInt32(&result, num*num)
		}(array[i])
	}

	wg.Wait()
	fmt.Println(result)
}

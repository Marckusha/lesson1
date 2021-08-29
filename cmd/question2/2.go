package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов
значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	var (
		array = []int{2, 4, 6, 8, 10}
		wg    sync.WaitGroup
	)

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(array[i])
	}
	wg.Wait()
}

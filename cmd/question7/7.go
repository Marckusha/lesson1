package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Реализовать конкурентную запись в map
*/

func main() {
	var (
		rangeRandNum = 3
		wg           sync.WaitGroup
		mapa         = map[int]int{}
		mu           sync.Mutex
	)

	rand.Seed(time.Now().UnixNano())

	//подсчитваем количество случайно-выпавших чисел в диапазоне от 0 до N
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				mu.Lock()
				mapa[rand.Intn(rangeRandNum)]++
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(mapa)
}

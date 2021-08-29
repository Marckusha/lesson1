package main

import (
	"fmt"
	"sync"
)

/*
Написать свою структуру счетчик, которая будет инкрементировать и выводить
значения в конкурентной среде
*/

type myStruct struct {
	value   int
	rwmutex sync.RWMutex
}

func (str *myStruct) inc() {
	str.rwmutex.Lock()
	str.value++
	str.rwmutex.Unlock()
}

func (str *myStruct) print() {
	str.rwmutex.RLock()
	fmt.Println(str.value)
	str.rwmutex.RUnlock()
}

func main() {
	str := myStruct{}

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(str *myStruct) {
			for i := 0; i < 100; i++ {
				str.inc()
				str.print()
			}
			wg.Done()
		}(&str)
	}
	wg.Wait()
}

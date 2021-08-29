package main

import (
	"fmt"
)

/*
Написать бинарный поиск встроенными методами языка
*/

func binSearch(arr []int, value int) (result int) {

	if len(arr) == 0 {
		return -1
	}

	mid := len(arr) / 2

	if arr[mid] > value {
		result = binSearch(arr[:mid], value)
	} else if arr[mid] < value {
		result = binSearch(arr[mid+1:], value)
		if result >= 0 {
			result += mid + 1
		}
	} else {
		result = mid
	}

	return
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 6, 7, 8}
	for i := 0; i < 10; i++ {
		fmt.Printf("index: %d value: %d\n", i, binSearch(arr, i))
	}
}

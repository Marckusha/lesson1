package main

import (
	"fmt"
)

/*
помещаем элементы отдельного массива в отдельную мапу
далее проверяем, чтобы элемент находился во всех мапах
тогда добавляем элемент в результирующий массив
*/

func getSet(arr []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, elem := range arr {
		set[elem] = struct{}{}
	}
	return set
}

func main() {
	arr1 := []int{8, 1, 5, 3, 2, 9, 6}
	arr2 := []int{9, 3, 7, 6, 9}
	var result []int

	set1 := getSet(arr1)
	set2 := getSet(arr2)

	for key := range set1 {
		if _, ok := set2[key]; ok {
			result = append(result, key)
		}
	}

	fmt.Println(result)
}

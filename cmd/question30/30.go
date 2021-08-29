package main

import "fmt"

/*Удалить i-ый элемент из слайса.*/

func remove(arr []int, n int) {
	if n < 0 || n >= len(arr) {
		return
	}

	arr = append(arr[0:n], arr[n+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	remove(slice, 4)
	fmt.Println(slice)
}

package question22

/*
Написать быструю сортировку встроенными методами языка
*/

func quickSort(arr []int, low int, high int) {
	if len(arr) == 0 {
		return
	}

	if low >= high {
		return
	}

	middle := low + (high-low)/2
	opora := arr[middle]
	i, j := low, high

	for i <= j {
		for ; arr[i] < opora; i++ {
		}
		for ; arr[j] > opora; j-- {
		}

		if i <= j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}

	if low < j {
		quickSort(arr, low, j)
	}

	if high > i {
		quickSort(arr, i, high)
	}
}

func main() {
}

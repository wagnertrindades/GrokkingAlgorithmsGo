package main

import "fmt"

func main() {
	my_list := []int{1, 3, 5, 7, 9, 10}

	fmt.Println(binary_search(my_list, 3))
	fmt.Println(binary_search(my_list, -1))
	fmt.Println(binary_search(my_list, 7))
}

func binary_search(list []int, item int) (int, bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		try := list[mid]

		if try == item {
			return mid, true
		}

		if try > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, false
}

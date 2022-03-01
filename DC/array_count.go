package main

import "fmt"

func main() {
	arr := []int{8, 5, 5, 6, 5, 9, 7, 8, 6, 2, 3}

	arrCount := arrCountRecursive(arr)

	fmt.Println(arrCount)
}

func arrCountRecursive(arr []int) int {

	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return 1
	} else {
		return 1 + arrCountRecursive(arr[1:])
	}
}

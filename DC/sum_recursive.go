package main

import "fmt"

func main() {
	arr := []int{2, 4, 6}

	sumArr := sum_recursive(arr)

	fmt.Println(sumArr)
}

func sum_recursive(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + sum_recursive(arr[1:])
	}
}

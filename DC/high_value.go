package main

import "fmt"

func main() {
	arr := []int{2, 9, 28, 8, 4, 25, 1, 50}

	minValue := -999999

	highValue := findHighValue(arr, minValue)

	fmt.Println(highValue)
}

func findHighValue(arr []int, highValue int) int {

	if len(arr) <= 0 {
		return 0
	} else if len(arr) == 1 {
		if arr[0] > highValue {
			return arr[0]
		}

		return highValue
	} else {
		if arr[0] > highValue {
			return findHighValue(arr[1:], arr[0])
		}

		return findHighValue(arr[1:], highValue)
	}
}

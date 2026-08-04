package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 7}

	fmt.Println(BinarySearch(x, 5))
}

func BinarySearch(x []int, t int) int {

	low := 0
	high := len(x) - 1

	for low <= high {
		mid := (low + high) / 2

		if x[mid] == t {
			return mid
		}

		if x[mid] < t {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

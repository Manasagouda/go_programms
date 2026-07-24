package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 7}

	fmt.Println(BinarySearch(x, 1))
}

func BinarySearch(x []int, t int) int {

	low, high := 0, len(x)-1

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

func MergeSortedArrays(x []int, max int) int {

	sum := max * (max + 1) / 2

	actual := 0

	for _, n := range x {
		actual = actual + n
	}

	return sum - actual
}

package main

import "fmt"

func main() {
	x := []int{1, 5, 2, 3, 8, 4, 9, 5, 7}

	fmt.Println(BubbleSort(x))
}

func BubbleSort(x []int) any {

	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-1-i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}

	}

	return x
}

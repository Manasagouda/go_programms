package main

import "fmt"

func main() {
	x := 123

	fmt.Println(ReverseOfNumber(x))
}

func ReverseOfNumber(x int) any {

	reverse := 0

	for x > 0 {
		reverse = reverse*10 + x%10

		x = x / 10
		fmt.Println(x)

	}

	return reverse
}

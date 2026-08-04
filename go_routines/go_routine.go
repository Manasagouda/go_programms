package main

import (
	"fmt"
	"time"
)

func main() {

	// single goroutine
	go PrintNumber()

	//multiple go routines
	go worker("B")
	go worker("A")

	time.Sleep(6 * time.Second)
}

func worker(s string) {

	for i := 0; i < 3; i++ {

		fmt.Println("letter with sequence", s, i)
		time.Sleep(time.Second)
	}
}

func PrintNumber() {
	for i := 0; i < 6; i++ {
		fmt.Println("number ", i)

		time.Sleep(time.Second)
	}
}

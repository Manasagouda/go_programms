package main

import (
	"fmt"
	"time"
)

func main() {

	// unbuffered channel
	ch := make(chan int)

	go square(ch, 5)

	fmt.Println("square of the number is", <-ch)

	time.Sleep(3 * time.Second)
}

func square(ch chan int, i int) {

	ch <- i * i
	time.Sleep(time.Second)
}

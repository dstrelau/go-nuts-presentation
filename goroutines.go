package main

import (
	"fmt"
	"time"
)

func receive(c chan int) {
	for n := range c {
		fmt.Println("received", n)
		time.Sleep(time.Second)
	}
}

func main() {
	queue := make(chan int, 5)

	receive(queue)

	for i := 0; i < 5; i++ {
		queue <- i
		fmt.Println("queued", i)
	}

	fmt.Println("done")
}

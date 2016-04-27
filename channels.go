package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	fmt.Println("queued one")
	queue <- "two"
	fmt.Println("queued two")

	fmt.Println("received", <-queue)
	fmt.Println("received", <-queue)
}

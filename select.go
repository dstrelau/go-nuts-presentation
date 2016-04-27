package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fill(c chan string, value string) {
	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		c <- value
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	foos := make(chan string)
	bars := make(chan string)

	go fill(foos, "foo")
	go fill(bars, "bar")
	timeout := time.After(10 * time.Second)

	now := time.Now()

	for {
		select {
		case foo := <-foos:
			fmt.Println(time.Since(now), foo)
		case bar := <-bars:
			fmt.Println(time.Since(now), bar)
		case <-timeout:
			fmt.Println("timeout!")
			return
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	theFuture := <-time.After(2 * time.Second)
	fmt.Println(theFuture)

	ticker := time.NewTicker(time.Second)
	fmt.Println(<-ticker.C)
	fmt.Println(<-ticker.C)
	fmt.Println(<-ticker.C)
}

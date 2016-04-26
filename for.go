package main

import "fmt"

func main() {
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}

	n := 1
	for n < 4 {
		fmt.Println(n)
		n++
	}

	a := []int{1, 2, 3}
	for i, n := range a {
		fmt.Println(i, n)
	}

	for {
		// ever
	}
}

package main

import "fmt"

func zero(n int) {
	n = 0
}

func zeroP(n *int) {
	*n = 0
}

func main() {
	n := 1

	zero(n) // copy value
	fmt.Println(n)

	zeroP(&n) // copy pointer
	fmt.Println(n)
}

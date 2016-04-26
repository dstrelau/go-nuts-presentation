package main

import (
	"fmt"
)

func main() {
	var array [3]int // fixed size
	fmt.Println("array:", array)
	fmt.Println("array first:", array[0])

	var slice []int // dynamic size
	fmt.Println("slice:", slice)
	slice = append(slice, 1)
	fmt.Println("slice first:", slice[0])

	parrots := []string{"party", "aussie", "bored", "conga", "shuffle"}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int // zero value
	fmt.Println(a)

	b := 1 // type inference
	fmt.Println(b)

	num, err := strconv.Atoi("2")
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}

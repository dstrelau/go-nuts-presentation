package main

import (
	"fmt"
	"strconv"
)

type fizzBuzz struct {
	Fizz bool
	Buzz bool
	Num  int
}

func (fb fizzBuzz) String() string {
	var s string

	if !fb.Fizz && !fb.Buzz {
		return strconv.Itoa(fb.Num)
	}
	if fb.Fizz {
		s += "fizz"
	}
	if fb.Buzz {
		s += "buzz"
	}

	return s
}

const howMany = 15

func main() {
	var fizzbuzzez [howMany]fizzBuzz

	for i := 0; i < howMany; i++ {
		f := (i+1)%3 == 0
		b := (i+1)%5 == 0
		fizzbuzzez[i] = fizzBuzz{Fizz: f, Buzz: b, Num: i + 1}
	}

	for _, fb := range fizzbuzzez {
		fmt.Println(fb)
	}
}

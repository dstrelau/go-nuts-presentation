package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// END TYPES OMIT
type shape interface {
	area() float64
}

func main() {
	c := circle{radius: 2}
	r := rectangle{width: 2, height: 3}
	shapes := [2]shape{c, r}
	for _, s := range shapes {
		fmt.Printf("%T %+v %.2f\n", s, s, s.area())
	}
}

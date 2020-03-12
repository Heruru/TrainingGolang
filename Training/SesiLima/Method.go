package SesiLima

import (
	"fmt"
	"math"
)

// method
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// method
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// method
func (x Test) Area() float64 {
	return x.a * x.b * x.c
}

func SesiLimaMethod() {
	c1 := Circle{10}
	c2 := Circle{25}
	r1 := Rectangle{9, 4}
	r2 := Rectangle{12, 2}
	x1 := Test{2, 3, 4}
	fmt.Println("Area Circle 1: ", c1.Area())
	fmt.Println("Area Circle 2: ", c2.Area())
	fmt.Println("Area Rectangle 1: ", r1.Area())
	fmt.Println("Area Rectangle 2: ", r2.Area())
	fmt.Println("Area Test 1: ", x1.Area())
}

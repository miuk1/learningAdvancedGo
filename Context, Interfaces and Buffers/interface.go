// Description: Interface example
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 2, Height: 3}
	fmt.Println(r.Area()) // prints 6

	var s Shape
	s = r
	fmt.Println(s.Area()) // also prints 6
}

package figure

import (
	"fmt"
	"math"
)

type Triangle struct {
	base   uint16
	height uint16
}

type Rectangle struct {
	width  uint16
	height uint16
}

type Circle struct {
	radius float64
}

type Squere struct {
	Side uint16
}

func (s *Squere) GetBeer() any {
	fmt.Println("Pam-param-pam-pam")
	return 1
}

func (t *Triangle) GetArea() (any, error) {
	if t.base > 0 && t.height > 0 {
		return (t.base * t.height) / 2, nil
	}
	return 0, fmt.Errorf("one of triangle input parameters is 0, base: %v height %v", t.base, t.height)
}

func (r *Rectangle) GetArea() (any, error) {
	if r.width > 0 && r.height > 0 {
		return r.width * r.height, nil
	}
	return 0, fmt.Errorf("one of rectangle input parameters is 0, width: %v height %v", r.width, r.height)
}

func (c *Circle) GetArea() (any, error) {
	if c.radius > 0 {
		return (c.radius * c.radius) * math.Pi, nil
	}
	return 0, fmt.Errorf("circle input parameter is 0, radius: %v", c.radius)
}

func NewTriangle(b uint16, h uint16) *Triangle {
	return &Triangle{
		base:   b,
		height: h,
	}
}

func NewRectangle(w uint16, h uint16) *Rectangle {
	return &Rectangle{
		width:  w,
		height: h,
	}
}

func NewCircle(r float64) *Circle {
	return &Circle{
		radius: r,
	}
}

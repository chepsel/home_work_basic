package figure

import (
	"fmt"
	"math"
)

type Triangle struct {
	base   float32
	height float32
}

type Rectangle struct {
	width  float32
	height float32
}

type Circle struct {
	radius float32
}

type Squere struct {
	Side uint16
}

func (s *Squere) GetBeer() any {
	fmt.Println("Pam-param-pam-pam")
	return 1
}

func (t *Triangle) GetArea() (float32, error) {
	if t.base > 0 && t.height > 0 {
		return (t.base * t.height) / 2, nil
	}
	return 0, fmt.Errorf("one of triangle input parameters is 0, base: %v height %v", t.base, t.height)
}

func (r *Rectangle) GetArea() (float32, error) {
	if r.width > 0 && r.height > 0 {
		return r.width * r.height, nil
	}
	return 0, fmt.Errorf("one of rectangle input parameters is 0, width: %v height %v", r.width, r.height)
}

func (c *Circle) GetArea() (float32, error) {
	if c.radius > 0 {
		return (c.radius * c.radius) * math.Pi, nil
	}
	return 0, fmt.Errorf("circle input parameter is 0, radius: %v", c.radius)
}

func NewTriangle(base float32, height float32) *Triangle {
	return &Triangle{
		base:   base,
		height: height,
	}
}

func NewRectangle(width float32, height float32) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func NewCircle(radius float32) *Circle {
	return &Circle{
		radius: radius,
	}
}

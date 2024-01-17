package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() (any, error)
}

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
	side uint16
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

func calculateArea(s any) (any, error) {
	switch t := s.(type) {
	case Shape:
		return s.(Shape).GetArea()
	default:
		return 0, fmt.Errorf("interface is not used, got: %T", &t)
	}
}

func main() {
	a := &Triangle{base: 1, height: 6}
	b := &Rectangle{width: 5, height: 6}
	c := &Circle{radius: 5}
	d := &Squere{side: 212}
	e := &Circle{radius: 0}

	if triangle, err := calculateArea(a); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("triangle area is: %v\n", triangle)
	}

	if rectangle, err := calculateArea(b); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("rectangle area is: %v\n", rectangle)
	}

	if circle, err := calculateArea(c); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("circle area is: %v\n", circle)
	}

	if squere, err := calculateArea(d); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("squere area is: %v\n", squere)
	}

	if zero, err := calculateArea(e); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("squere area is: %v\n", zero)
	}
}

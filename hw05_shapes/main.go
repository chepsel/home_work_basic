package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw05_shapes/figure"
)

type Shape interface {
	GetArea() (float32, error)
}

type Kukareku interface {
	GetBeer() any
}

func calculateArea(s any) (float32, error) {
	switch t := s.(type) {
	case Shape:
		return t.GetArea()
	default:
		return 0, fmt.Errorf("interface is not used, got: %T", &t)
	}
}

func main() {
	var a Shape = figure.NewTriangle(22, 98)
	var b Shape = figure.NewRectangle(3, 44)
	var c Shape = figure.NewCircle(8)
	var d Kukareku = &figure.Squere{Side: 212}
	var e Shape = figure.NewCircle(0)
	var f Shape = figure.NewRectangle(0, 1)

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

	if zero, err := calculateArea(f); err != nil {
		fmt.Printf("unable to calculate area: %v\n", err)
	} else {
		fmt.Printf("squere area is: %v\n", zero)
	}
}

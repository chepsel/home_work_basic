package shapes

import (
	"errors"
	"fmt"

	"github.com/chepsel/home_work_basic/hw06_testing/shapes/figure"
)

type shape interface {
	GetArea() (float32, error)
}

type figureType uint8

const (
	Triangle figureType = iota
	Rectangle
	Circle
	Squere
)

func (f figureType) ObjectParams() int {
	return [...]int{2, 2, 1, 1}[f]
}

var (
	ErrWrongParamsNum = errors.New("wrong number of params")
	ErrWrongValue     = errors.New("wrong value")
)

func calculateArea(s any) (float32, error) {
	switch t := s.(type) {
	case shape:
		return t.GetArea()
	default:
		return 0, fmt.Errorf("shape is not used, got: %T", &t)
	}
}

func (f figureType) validateParams(params ...float32) error {
	paramsCount := len(params)
	switch {
	case paramsCount != f.ObjectParams():
		return ErrWrongParamsNum
	default:
		return nil
	}
}

func (f figureType) calculate(params ...float32) (float32, error) {
	switch f {
	case Triangle:
		return calculateArea(figure.NewTriangle(params[0], params[1]))
	case Rectangle:
		return calculateArea(figure.NewRectangle(params[0], params[1]))
	case Circle:
		return calculateArea(figure.NewCircle(params[0]))
	case Squere:
		return calculateArea(&figure.Squere{Side: uint16(params[0])})
	default:
		return 0, ErrWrongValue
	}
}

func (f figureType) FigureArea(params ...float32) (float32, error) {
	if err := f.validateParams(params...); err != nil {
		return 0, err
	}
	return f.calculate(params...)
}

package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFigureAreaErr(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		want   error
		desc   string
		input1 figureType
		input2 float32
		input3 float32
		input4 float32
	}{
		{
			desc:   "1 element in triangle",
			input1: Triangle,
			input2: 1.3,
			want:   ErrWrongParamsNum,
		},
		{
			desc:   "1 element in rectangle",
			input1: Rectangle,
			input2: 5,
			want:   ErrWrongParamsNum,
		},
		{
			desc:   "2 element in circle",
			input1: Circle,
			input2: 5,
			input3: 5,
			want:   ErrWrongParamsNum,
		},
		{
			desc:   "3 element in triangle",
			input1: Triangle,
			input2: 1.3,
			input3: 4,
			input4: 5,
			want:   ErrWrongParamsNum,
		},
		{
			desc:   "3 element in rectangle",
			input1: Rectangle,
			input2: 1.3,
			input3: 4,
			input4: 5,
			want:   ErrWrongParamsNum,
		},
		{
			desc:   "3 element in circle",
			input1: Circle,
			input2: 1.3,
			input3: 4,
			input4: 5,
			want:   ErrWrongParamsNum,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			switch {
			case tC.input3 == 0 && tC.input4 == 0:
				_, err := tC.input1.FigureArea(tC.input2)
				assert.Equal(t, tC.want, err)
			case (tC.input3 != 0 && tC.input4 == 0):
				_, err := tC.input1.FigureArea(tC.input2, tC.input3)
				assert.Equal(t, tC.want, err)
			case tC.input3 != 0 && tC.input4 != 0:
				_, err := tC.input1.FigureArea(tC.input2, tC.input3, tC.input4)
				assert.Equal(t, tC.want, err)
			default:
				t.Errorf("somthing wrong")
			}
		})
	}
}

func TestFigureArea(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		want   float32
		desc   string
		input1 figureType
		input2 float32
		input3 float32
	}{
		{
			desc:   "2 element in triangle",
			input1: Triangle,
			input2: 1.3,
			input3: 4,
			want:   2.6,
		},
		{
			desc:   "2 element in rectangle",
			input1: Rectangle,
			input2: 1.3,
			input3: 4,
			want:   5.2,
		},
		{
			desc:   "1 element in circle",
			input1: Circle,
			input2: 1.3,
			want:   5.3092914,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input1 == Circle {
				got, err := tC.input1.FigureArea(tC.input2)
				if err != nil {
					t.Errorf("somthing wrong")
				}
				assert.Equal(t, tC.want, got)
			} else {
				got, err := tC.input1.FigureArea(tC.input2, tC.input3)
				if err != nil {
					t.Errorf("somthing wrong")
				}
				assert.Equal(t, tC.want, got)
			}
		})
	}
}

func TestFigureAreaCheckInf(t *testing.T) {
	_, err := Squere.FigureArea(1)
	if err == nil {
		t.Errorf("missing error")
	}
}

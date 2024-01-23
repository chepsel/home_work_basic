package figure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTriangle(t *testing.T) {
	testCases := []struct {
		want *Triangle
		desc string
		a, b float32
	}{
		{
			desc: "Triangle 1,2",
			a:    1,
			b:    2,
			want: &Triangle{
				base:   1,
				height: 2,
			},
		},
		{
			desc: "Triangle 0,2",
			a:    0,
			b:    2,
			want: &Triangle{
				base:   0,
				height: 2,
			},
		},
		{
			desc: "Triangle 1.1,9.8",
			a:    1.1,
			b:    9.8,
			want: &Triangle{
				base:   1.1,
				height: 9.8,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := NewTriangle(tC.a, tC.b)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestNewRectangle(t *testing.T) {
	testCases := []struct {
		want *Rectangle
		desc string
		a, b float32
	}{
		{
			desc: "Rectangle 1,2",
			a:    1,
			b:    2,
			want: &Rectangle{
				width:  1,
				height: 2,
			},
		},
		{
			desc: "Rectangle 0,2",
			a:    0,
			b:    2,
			want: &Rectangle{
				width:  0,
				height: 2,
			},
		},
		{
			desc: "Rectangle 1.1,9.8",
			a:    1.1,
			b:    9.8,
			want: &Rectangle{
				width:  1.1,
				height: 9.8,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := NewRectangle(tC.a, tC.b)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestNewCircle(t *testing.T) {
	testCases := []struct {
		want *Circle
		desc string
		a    float32
	}{
		{
			desc: "Circle 1",
			a:    1,
			want: &Circle{
				radius: 1,
			},
		},
		{
			desc: "Circle 0",
			a:    0,
			want: &Circle{
				radius: 0,
			},
		},
		{
			desc: "Circle 5.3",
			a:    5.3,
			want: &Circle{
				radius: 5.3,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := NewCircle(tC.a)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestGetAreaRectangleErr(t *testing.T) {
	input := &Rectangle{
		width:  0,
		height: 1,
	}
	_, err := input.GetArea()
	if err == nil {
		t.Errorf("missing error")
	}
}

func TestGetAreaTriangle(t *testing.T) {
	input := &Triangle{
		base:   6.2,
		height: 8.5,
	}
	var want float32 = 26.349998
	got, err := input.GetArea()
	if err != nil {
		t.Errorf("some error")
	}
	assert.Equal(t, want, got)
}

func TestGetAreaTriangleErr(t *testing.T) {
	input := &Triangle{
		base:   0,
		height: 1,
	}
	_, err := input.GetArea()
	if err == nil {
		t.Errorf("missing error")
	}
}

func TestGetAreaRectangle(t *testing.T) {
	input := &Rectangle{
		width:  5,
		height: 8.5,
	}
	var want float32 = 42.5
	got, err := input.GetArea()
	if err != nil {
		t.Errorf("some error")
	}
	assert.Equal(t, want, got)
}

func TestGetAreaCircleErr(t *testing.T) {
	input := &Circle{
		radius: 0,
	}
	_, err := input.GetArea()
	if err == nil {
		t.Errorf("missing error")
	}
}

func TestGetAreaCircle(t *testing.T) {
	input := &Circle{
		radius: 5,
	}
	var want float32 = 78.53982
	got, err := input.GetArea()
	if err != nil {
		t.Errorf("some error")
	}
	assert.Equal(t, want, got)
}

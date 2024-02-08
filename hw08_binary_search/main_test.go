package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func typeofObject(variable interface{}) string {
	switch variable.(type) {
	case bool:
		return "boolean"
	default:
		return "other"
	}
}

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		want   any
		desc   string
		input1 []int
		input2 int
	}{
		{
			desc:   "check array(true)",
			input1: []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 4,
			want:   true,
		},
		{
			desc:   "check array(false)",
			input1: []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 2,
			want:   false,
		},
		{
			desc:   "check array(digit - one step)",
			input1: []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 8,
			want:   true,
		},
		{
			desc:   "check array(digit - wrong sorting)",
			input1: []int{1, 4, 5, 6, 7, 9, 10, 11, 12, 8},
			input2: 8,
			want:   false,
		},
		{
			desc:   "check array(4)",
			input1: []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 4,
			want:   4,
		},
		{
			desc:   "check array(0)",
			input1: []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 0,
			want:   0,
		},
		{
			desc:   "check array(8)",
			input1: []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			input2: 8,
			want:   8,
		},
		{
			desc:   "check array(!=8)",
			input1: []int{1, 4, 5, 6, 7, 9, 10, 11, 12, 8},
			input2: 8,
			want:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if typeofObject(tC.want) == "boolean" {
				_, got := binarySearch(&tC.input1, tC.input2)
				assert.Equal(t, tC.want, got)
			} else {
				got, _ := binarySearch(&tC.input1, tC.input2)
				assert.Equal(t, tC.want, got)
			}
		})
	}
}

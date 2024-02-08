package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeString(t *testing.T) {
	testCases := []struct {
		want   any
		desc   string
		input1 string
		input2 bool
	}{
		{
			desc:   "check array(1)",
			input1: "раз два: три, четыре пять шесть, \r  string int rune, раз",
			input2: true,
			want: map[string]int{
				"int":    1,
				"rune":   1,
				"string": 1,
				"два":    1,
				"пять":   1,
				"раз":    2,
				"три":    1,
				"четыре": 1,
				"шесть":  1,
			},
		},
		{
			desc:   "check array(2)",
			input1: "Как дела?",
			input2: true,
			want: map[string]int{
				"дела": 1,
				"как":  1,
			},
		},
		{
			desc:   "check error(1)",
			input1: " , . \n \r ?",
			input2: false,
			want:   ErrWrongValue,
		},
		{
			desc:   "check error(2)",
			input1: "",
			input2: false,
			want:   ErrWrongValue,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input2 == true {
				got, _ := makeString(tC.input1)
				assert.Equal(t, tC.want, got)
			} else {
				_, got := makeString(tC.input1)
				assert.Equal(t, tC.want, got)
			}
		})
	}
}

func TestValidateString(t *testing.T) {
	testCases := []struct {
		want   any
		desc   string
		input1 string
	}{
		{
			desc:   "check word(1)",
			input1: "слово",
			want:   "слово",
		},
		{
			desc:   "check word(2)",
			input1: "слово,;",
			want:   "слово\x00\x00",
		},
		{
			desc:   "check word(3)",
			input1: " ",
			want:   "\x00",
		},
		{
			desc:   "check emty(1)",
			input1: "",
			want:   "",
		},
		{
			desc:   "check digit(1)",
			input1: "123",
			want:   "\x00\x00\x00",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := validateString(tC.input1)
			assert.Equal(t, tC.want, got)
		})
	}
}

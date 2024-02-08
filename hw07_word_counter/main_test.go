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
			want:   []string{"раз", "два", "три", "четыре", "пять", "шесть", "string", "int", "rune", "раз"},
		},
		{
			desc:   "check array(2)",
			input1: "Как дела?",
			input2: true,
			want:   []string{"как", "дела"},
		},
		{
			desc:   "check error(1)",
			input1: " , . \n \r ?",
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

func TestCountWords(t *testing.T) {
	testCases := []struct {
		want      map[string]int
		desc      string
		testError bool
		input1    string
		wantErr   error
	}{
		{
			desc:      "check words(1)",
			input1:    "слово dsd слово",
			testError: false,
			want:      map[string]int{"dsd": 1, "слово": 2},
		},
		{
			desc:      "check emty(1)",
			input1:    "",
			testError: false,
			want:      nil,
		},
		{
			desc:      "check words(2)",
			input1:    "раз два три",
			testError: false,
			want:      map[string]int{"раз": 1, "два": 1, "три": 1},
		},
		{
			desc:      "check error(1)",
			input1:    "",
			testError: true,
			wantErr:   ErrWrongValue,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.testError == false {
				got, _ := countWords(tC.input1)
				assert.Equal(t, tC.want, got)
			} else if tC.testError == true {
				_, err := countWords(tC.input1)
				assert.Equal(t, tC.wantErr, err)
			}
		})
	}
}

package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateMakeLine(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		want, desc     string
		input1, input2 int
	}{
		{
			desc:   "9 elements, starts 1",
			input1: 9,
			input2: 1,
			want:   " # # # #\n",
		},
		{
			desc:   "3 elements, starts 0",
			input1: 3,
			input2: 0,
			want:   "# #\n",
		},
		{
			desc:   "22 elements, starts 0",
			input1: 22,
			input2: 0,
			want:   "# # # # # # # # # # # \n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := makeLine(tC.input1, tC.input2)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestChessboard(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		want, desc string
		input1     int
	}{
		{
			desc:   "9 elements",
			input1: 9,
			want:   "# # # # #\n # # # # \n# # # # #\n # # # # \n# # # # #\n # # # # \n# # # # #\n # # # # \n# # # # #\n",
		},
		{
			desc:   "3 elements",
			input1: 3,
			want:   "# #\n # \n# #\n",
		},
		{
			desc:   "5 elements",
			input1: 5,
			want:   "# # #\n # # \n# # #\n # # \n# # #\n",
		},
		{
			desc:   "0 elements",
			input1: 0,
			want:   "# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := Chessboard(tC.input1)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestEmptyInput(t *testing.T) {
	_, err := getInput()
	if err == nil {
		t.Errorf("missing error")
	}
}

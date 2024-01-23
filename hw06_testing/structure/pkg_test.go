package structure

import (
	"testing"

	"github.com/chepsel/home_work_basic/hw06_testing/structure/books"
	"github.com/stretchr/testify/assert"
)

func TestNewComparator(t *testing.T) {
	testCases := []struct {
		want  *Comparator
		desc  string
		input compareField
	}{
		{
			desc:  "Check year",
			input: Year,
			want:  &Comparator{fieldCompare: Year},
		},
		{
			desc:  "Check size",
			input: Size,
			want:  &Comparator{fieldCompare: Size},
		},
		{
			desc:  "Check rate",
			input: Rate,
			want:  &Comparator{fieldCompare: Rate},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := NewComparator(tC.input)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestNewBook(t *testing.T) {
	testCases := []struct {
		want   bool
		desc   string
		input1 *books.Book
		input2 *books.Book
		input3 *Comparator
	}{
		{
			desc:   "Year Check 1",
			input1: NewBook("", "", "", 2019, 1, 1),
			input2: NewBook("", "", "", 2009, 1, 1),
			input3: NewComparator(Year),
			want:   true,
		},
		{
			desc:   "Year Check 2",
			input1: NewBook("", "", "", 1019, 1, 1),
			input2: NewBook("", "", "", 2009, 1, 1),
			input3: NewComparator(Year),
			want:   false,
		},
		{
			desc:   "Size Check 1",
			input1: NewBook("", "", "", 2009, 400, 2.4),
			input2: NewBook("", "", "", 2009, 100, 2.4),
			input3: NewComparator(Size),
			want:   true,
		},
		{
			desc:   "Size Check 2",
			input1: NewBook("", "", "", 2009, 100, 2.4),
			input2: NewBook("", "", "", 2009, 400, 2.4),
			input3: NewComparator(Size),
			want:   false,
		},
		{
			desc:   "Rate Check 1",
			input1: NewBook("", "", "", 2009, 100, 3.4),
			input2: NewBook("", "", "", 2009, 400, 2.4),
			input3: NewComparator(Rate),
			want:   true,
		},
		{
			desc:   "Rate Check 2",
			input1: NewBook("", "", "", 2009, 100, 1.4),
			input2: NewBook("", "", "", 2009, 400, 2.4),
			input3: NewComparator(Rate),
			want:   false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.input3.Compare(tC.input1, tC.input2)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestFieldName(t *testing.T) {
	testCases := []struct {
		want  string
		desc  string
		input compareField
	}{
		{
			desc:  "Check year",
			input: Year,
			want:  "year",
		},
		{
			desc:  "Check size",
			input: Size,
			want:  "size",
		},
		{
			desc:  "Check rate",
			input: Rate,
			want:  "rate",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.input.String()
			assert.Equal(t, tC.want, got)
		})
	}
}

package books

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		want      *Book
		desc      string
		input1    []byte
		input2    *Book
		testError bool
	}{
		{
			desc:   "check valid",
			input1: []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400,"rate":2.4}`),
			input2: &Book{},
			want: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			testError: false,
		},
		{
			desc:   "check missing rate",
			input1: []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400}`),
			input2: &Book{},
			want: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			testError: false,
		},
		{
			desc:   "check missing size",
			input1: []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"rate":2.4}`),
			input2: &Book{},
			want: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Rate:   2.4,
			},
			testError: false,
		},
		{
			desc:   "check missing author",
			input1: []byte(`{"id":"978","title":"Элегантность ежика","year":2009,"rate":2.4}`),
			input2: &Book{},
			want: &Book{
				ID:    "978",
				Title: "Элегантность ежика",
				Year:  2009,
				Rate:  2.4,
			},
			testError: false,
		},
		{
			desc:      "check wrong format",
			input1:    []byte(`{"id":"978","title":"Элегантность ежика,"year":2009,"rate":2.4}`),
			input2:    &Book{},
			want:      &Book{},
			testError: true,
		},
		{
			desc:      "check missing comma",
			input1:    []byte(`{"id":"978""title":"Элегантность ежика,"year":2009,"rate":2.4}`),
			input2:    &Book{},
			want:      &Book{},
			testError: true,
		},
		{
			desc:      "check missing value",
			input1:    []byte(`{"id":"978","title":"Элегантность ежика,"year":2009,"rate":`),
			input2:    &Book{},
			want:      &Book{},
			testError: true,
		},
		{
			desc:      "check missing all",
			input1:    []byte(``),
			input2:    &Book{},
			want:      &Book{},
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.testError {
				err := tC.input2.UnmarshalJSON(tC.input1)
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				tC.input2.UnmarshalJSON(tC.input1)
				assert.Equal(t, tC.want, tC.input2)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		want   []byte
		desc   string
		input1 *Book
		input2 *Book
	}{
		{
			desc:   "check valid",
			want:   []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400,"rate":2.4}`),
			input2: &Book{},
			input1: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
		},
		{
			desc:   "check missing rate",
			want:   []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400}`),
			input2: &Book{},
			input1: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
		},
		{
			desc:   "check missing size",
			want:   []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"rate":2.4}`),
			input2: &Book{},
			input1: &Book{
				ID:     "978",
				Title:  "ежик",
				Author: "Мюриель Барбери",
				Year:   2009,
				Rate:   2.4,
			},
		},
		{
			desc:   "check missing author",
			want:   []byte(`{"id":"978","title":"Элегантность ежика","year":2009,"rate":2.4}`),
			input2: &Book{},
			input1: &Book{
				ID:    "978",
				Title: "Элегантность ежика",
				Year:  2009,
				Rate:  2.4,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := tC.input1.MarshalJSON()
			assert.Equal(t, string(tC.want), string(got))
		})
	}
}

func TestNewBook(t *testing.T) {
	const id, title, author string = "1", "2", "3"
	const year, size uint32 = 3, 2
	const rate float32 = 2.2
	want := &Book{ID: id, Title: title, Author: author, Year: year, Size: size, Rate: rate}
	got := NewBook(id, title, author, year, size, rate)
	assert.Equal(t, want, got)
}

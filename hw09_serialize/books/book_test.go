package books

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	testCases := []struct {
		want      JSONBook
		desc      string
		input1    []byte
		input2    []byte
		input3    JSONBook
		testError bool
	}{
		{
			desc:   "check one",
			input1: []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400,"rate":2.4}`),
			input3: JSONBook{},
			want: JSONBook{
				Books: []Book{
					{
						ID:     "978",
						Title:  "ежик",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
						Rate:   2.4,
					},
				},
			},
			testError: false,
		},
		{
			desc:   "check many",
			input1: []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"rate":2.4}`),
			input2: []byte(`,{"id":"2","title":"б","author":"МБ","year":7,"size":30}]`),
			input3: JSONBook{},
			want: JSONBook{
				Books: []Book{
					{
						ID:     "1",
						Title:  "е",
						Author: "МБ",
						Year:   2,
						Rate:   2.4,
					}, {
						ID:     "2",
						Title:  "б",
						Author: "МБ",
						Year:   7,
						Size:   30,
					},
				},
			},
			testError: false,
		},
		{
			desc:   "check wrong format",
			input1: []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"size":40,"rate":2.4}`),
			input2: []byte(`,{"id":"2","title":"б}{"id":"2","title":"б}]`),
			input3: JSONBook{},
			want: JSONBook{
				Books: []Book{},
			},
			testError: true,
		},
		{
			desc:   "check missing comma",
			input1: []byte(`{"id":"1","title":"е","author":"МБ","year":2,"size":40"rate":2.4}`),
			input3: JSONBook{},
			want: JSONBook{
				Books: []Book{},
			},
			testError: true,
		},
		{
			desc:   "check regular string",
			input1: []byte(`asad`),
			input3: JSONBook{},
			want: JSONBook{
				Books: []Book{},
			},
			testError: true,
		},
		{
			desc:      "check missing all",
			input1:    []byte(``),
			input3:    JSONBook{},
			want:      JSONBook{},
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.input1 = append(tC.input1, tC.input2...)
			err := tC.input3.UnmarshalJSON(tC.input1)
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				assert.Equal(t, tC.want, tC.input3)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		want1     []byte
		want2     []byte
		desc      string
		input1    *JSONBook
		input2    *JSONBook
		testError bool
	}{
		{
			desc:   "check valid",
			want1:  []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"size":40,"rate":2.4}`),
			want2:  []byte(`,{"id":"2","title":"б","author":"МБ","year":7,"size":30}]`),
			input2: &JSONBook{},
			input1: &JSONBook{
				Books: []Book{
					{
						ID:     "1",
						Title:  "е",
						Author: "МБ",
						Year:   2,
						Size:   40,
						Rate:   2.4,
					}, {
						ID:     "2",
						Title:  "б",
						Author: "МБ",
						Year:   7,
						Size:   30,
					},
				},
			},
			testError: false,
		},
		{
			desc:   "check one",
			want1:  []byte(`{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400}`),
			input2: &JSONBook{},
			input1: &JSONBook{
				Books: []Book{
					{
						ID:     "978",
						Title:  "ежик",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
					},
				},
			},
			testError: false,
		},
		{
			desc:   "check missing size",
			want1:  []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"size":40,"rate":2.4}`),
			want2:  []byte(`,{"id":"2","title":"б","year":0}]`),
			input2: &JSONBook{},
			input1: &JSONBook{
				Books: []Book{
					{
						ID:     "1",
						Title:  "е",
						Author: "МБ",
						Year:   2,
						Size:   40,
						Rate:   2.4,
					}, {
						ID:    "2",
						Title: "б",
					},
				},
			},
			testError: false,
		},
		{
			desc:      "check missing all",
			want1:     []byte(``),
			input2:    &JSONBook{},
			input1:    &JSONBook{},
			testError: false,
		},
		{
			desc:      "check missing all",
			want2:     []byte(``),
			input2:    &JSONBook{},
			input1:    &JSONBook{},
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.input1.MarshalJSON()
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				tC.want1 = append(tC.want1, tC.want2...)
				assert.Equal(t, string(tC.want1), string(got))
			}
		})
	}
}

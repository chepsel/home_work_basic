package main

import (
	"testing"

	"github.com/chepsel/home_work_basic/hw09_serialize/books"
	"github.com/chepsel/home_work_basic/hw09_serialize/protoc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestToProtobufStructure(t *testing.T) {
	testCases := []struct {
		want      *protoc.BooksSlice
		desc      string
		input1    []*books.Book
		testError bool
	}{
		{
			desc: "check valid",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
					Rate:   2.4,
				}, {
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
					Rate:   2.4,
				},
			},
			want: &protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
						Rate:   2.4,
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
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
			desc: "check missing rate",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
				},
			},
			want: &protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
					},
				},
			},
			testError: false,
		},
		{
			desc: "check missing size and rate",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Rate:   2.4,
				}, {
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
				},
			},
			want: &protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Rate:   2.4,
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
					},
				},
			},
			testError: false,
		},
		{
			desc: "check one element",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
				},
			},
			want: &protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ToProtobufStructure(tC.input1)
			want := proto.Clone(tC.want).(*protoc.BooksSlice)
			assert.Equal(t, want, got)
		})
	}
}

const (
	protoShit1 = "\n\x11978-5-389-21499-6\x12"
	protoShit2 = "#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x035\x9a\x99\x19@"
	protoShit3 = "\v ежика\x1a\x1dМюриель Барбери \xd9\x0f"
	protoShit4 = "\"Элегантность ежикa\x1a\x1dМюриель Барбери \xd9\x0f"
	protoShit5 = "#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x03"
)

func TestUnmarshalPBSlice(t *testing.T) {
	testCases := []struct {
		want      protoc.BooksSlice
		desc      string
		input1    []*books.Book
		input2    []byte
		testError bool
	}{
		{
			desc:   "check valid",
			input1: []*books.Book{},
			want: protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
						Rate:   2.4,
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
						Rate:   2.4,
					},
				},
			},
			input2:    []byte("\nb" + protoShit1 + protoShit2 + "\nb" + protoShit1 + protoShit2),
			testError: false,
		},
		{
			desc:   "check missing rate and size",
			input1: []*books.Book{},
			want: protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  " ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежикa",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			input2:    []byte("\nB" + protoShit1 + protoShit3 + "\nY" + protoShit1 + protoShit4),
			testError: false,
		},
		{
			desc:   "check one",
			input1: []*books.Book{},
			want: protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  " ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			input2:    []byte("\nB" + protoShit1 + protoShit3),
			testError: false,
		},
		{
			desc:   "check error",
			input1: []*books.Book{},
			want: protoc.BooksSlice{
				Books: []*protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  " ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежикa",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			input2:    []byte("sfdsdfsdfsdf"),
			testError: true,
		},
	}
	for tI := range testCases {
		t.Run(testCases[tI].desc, func(t *testing.T) {
			protoBookNew, err := UnmarshalPBSlice(testCases[tI].input2)
			// fmt.Println(protoBookNew)
			if testCases[tI].testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				if err != nil {
					t.Errorf("wrong protostructure error")
				}
				for i := range protoBookNew.Books {
					element := proto.Clone(protoBookNew.Books[i]).(*protoc.Book)
					assert.Equal(t, testCases[tI].want.Books[i].GetID(), element.GetID())
					assert.Equal(t, testCases[tI].want.Books[i].GetAuthor(), element.GetAuthor())
					assert.Equal(t, testCases[tI].want.Books[i].GetRate(), element.GetRate())
					assert.Equal(t, testCases[tI].want.Books[i].GetSize(), element.GetSize())
					assert.Equal(t, testCases[tI].want.Books[i].GetYear(), element.GetYear())
					assert.Equal(t, testCases[tI].want.Books[i].GetTitle(), element.GetTitle())
					// fmt.Printf("\nleft:%s right:%s\n", tC.want.Books[i].GetID(), element.GetID())
				}
			}
		})
	}
}

func TestUnmarshalPB(t *testing.T) {
	testCases := []struct {
		want      *protoc.Book
		desc      string
		input1    []*books.Book
		input2    []byte
		testError bool
	}{
		{
			desc:   "check valid",
			input1: []*books.Book{},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			input2:    []byte(protoShit1 + protoShit2),
			testError: false,
		},
		{
			desc:   "check missing rate and size",
			input1: []*books.Book{},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  " ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			input2:    []byte(protoShit1 + protoShit3),
			testError: false,
		},
		{
			desc:   "check one",
			input1: []*books.Book{},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  " ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			input2:    []byte(protoShit1 + protoShit3),
			testError: false,
		},
		{
			desc:   "check error",
			input1: []*books.Book{},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  " ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			input2:    []byte("sfdsdfsdfsdf"),
			testError: true,
		},
	}
	for tI := range testCases {
		t.Run(testCases[tI].desc, func(t *testing.T) {
			protoBookNew, err := UnmarshalPB(testCases[tI].input2)
			// fmt.Println(protoBookNew)
			if testCases[tI].testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				if err != nil {
					t.Errorf("wrong protostructure error")
				}
				element := proto.Clone(protoBookNew).(*protoc.Book)
				assert.Equal(t, testCases[tI].want.GetID(), element.GetID())
				assert.Equal(t, testCases[tI].want.GetAuthor(), element.GetAuthor())
				assert.Equal(t, testCases[tI].want.GetRate(), element.GetRate())
				assert.Equal(t, testCases[tI].want.GetSize(), element.GetSize())
				assert.Equal(t, testCases[tI].want.GetYear(), element.GetYear())
				assert.Equal(t, testCases[tI].want.GetTitle(), element.GetTitle())
				// fmt.Printf("\nleft:%s right:%s\n", tC.want.Books[i].GetID(), element.GetID())
			}
		})
	}
}

func TestMarshalPBSlice(t *testing.T) {
	testCases := []struct {
		input1    []*books.Book
		desc      string
		want      []byte
		testError bool
	}{
		{
			desc: "check valid",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
					Rate:   2.4,
				}, {
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
					Rate:   2.4,
				},
			},
			want:      []byte("\nb" + protoShit1 + protoShit2 + "\nb" + protoShit1 + protoShit2),
			testError: false,
		},
		{
			desc: "check missing rate",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
				}, {
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
				},
			},
			want:      []byte("\n]" + protoShit1 + protoShit5 + "\n]" + protoShit1 + protoShit5),
			testError: false,
		},
		{
			desc: "check missing size and rate",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  " ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
				}, {
					ID:     "978-5-389-21499-6",
					Title:  "Элегантность ежикa",
					Author: "Мюриель Барбери",
					Year:   2009,
				},
			},
			want:      []byte("\nB" + protoShit1 + protoShit3 + "\nY" + protoShit1 + protoShit4),
			testError: false,
		},
		{
			desc: "check one",
			input1: []*books.Book{
				{
					ID:     "978-5-389-21499-6",
					Title:  " ежика",
					Author: "Мюриель Барбери",
					Year:   2009,
				},
			},
			want:      []byte("\nB" + protoShit1 + protoShit3),
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tmp := ToProtobufStructure(tC.input1)
			got, err := MarshalPBSlice(tmp)
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				assert.Equal(t, string(tC.want), string(got))
			}
		})
	}
}

func TestMarshalPB(t *testing.T) {
	testCases := []struct {
		input1    *protoc.Book
		desc      string
		want      []byte
		testError bool
	}{
		{
			desc: "check valid",
			input1: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			want:      []byte(protoShit1 + protoShit2),
			testError: false,
		},
		{
			desc: "check missing rate",
			input1: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			want:      []byte(protoShit1 + protoShit5),
			testError: false,
		},
		{
			desc: "check missing size and rate",
			input1: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  " ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			want:      []byte(protoShit1 + protoShit3),
			testError: false,
		},
		{
			desc: "check one",
			input1: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  " ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			want:      []byte(protoShit1 + protoShit3),
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := MarshalPB(tC.input1)
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				assert.Equal(t, string(tC.want), string(got))
			}
		})
	}
}

func TestMarshalJSONSlice(t *testing.T) {
	testCases := []struct {
		want1     []byte
		want2     []byte
		desc      string
		input1    []*books.Book
		input2    []*books.Book
		testError bool
	}{
		{
			desc:   "check valid",
			want1:  []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"size":40,"rate":2.4}`),
			want2:  []byte(`,{"id":"2","title":"б","author":"МБ","year":7,"size":30}]`),
			input2: []*books.Book{},
			input1: []*books.Book{
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
			testError: false,
		},
		{
			desc:   "check one",
			want1:  []byte(`[{"id":"978","title":"ежик","author":"Мюриель Барбери","year":2009,"size":400}]`),
			input2: []*books.Book{},
			input1: []*books.Book{
				{
					ID:     "978",
					Title:  "ежик",
					Author: "Мюриель Барбери",
					Year:   2009,
					Size:   400,
				},
			},
			testError: false,
		},
		{
			desc:   "check missing size",
			want1:  []byte(`[{"id":"1","title":"е","author":"МБ","year":2,"size":40,"rate":2.4}`),
			want2:  []byte(`,{"id":"2","title":"б","year":0}]`),
			input2: []*books.Book{},
			input1: []*books.Book{
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
			testError: false,
		},
		{
			desc:      "check missing all",
			want1:     []byte(`[]`),
			input2:    []*books.Book{},
			input1:    []*books.Book{},
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := MarshalJSONSlice(tC.input1)
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

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
		input1    *books.JSONBook
		testError bool
	}{
		{
			desc: "check valid",
			input1: &books.JSONBook{
				Books: []books.Book{
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
			input1: &books.JSONBook{
				Books: []books.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
					},
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
			input1: &books.JSONBook{
				Books: []books.Book{
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
			input1: &books.JSONBook{
				Books: []books.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
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

func TestUnmarshalProto(t *testing.T) {
	testCases := []struct {
		want      protoc.BooksSlice
		desc      string
		input1    *books.JSONBook
		input2    []byte
		testError bool
	}{
		{
			desc:   "check valid",
			input1: &books.JSONBook{},
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
			input1: &books.JSONBook{},
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
			input1: &books.JSONBook{},
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
			input1: &books.JSONBook{},
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
			protoBookNew, err := UnmarshalProto(testCases[tI].input2)
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

func TestMarshalProto(t *testing.T) {
	testCases := []struct {
		input1    *books.JSONBook
		desc      string
		want      []byte
		testError bool
	}{
		{
			desc: "check valid",
			input1: &books.JSONBook{
				Books: []books.Book{
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
			want:      []byte("\nb" + protoShit1 + protoShit2 + "\nb" + protoShit1 + protoShit2),
			testError: false,
		},
		{
			desc: "check missing rate",
			input1: &books.JSONBook{
				Books: []books.Book{
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
			},
			want:      []byte("\n]" + protoShit1 + protoShit5 + "\n]" + protoShit1 + protoShit5),
			testError: false,
		},
		{
			desc: "check missing size and rate",
			input1: &books.JSONBook{
				Books: []books.Book{
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
			want:      []byte("\nB" + protoShit1 + protoShit3 + "\nY" + protoShit1 + protoShit4),
			testError: false,
		},
		{
			desc: "check one",
			input1: &books.JSONBook{
				Books: []books.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  " ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			want:      []byte("\nB" + protoShit1 + protoShit3),
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tmp := ToProtobufStructure(tC.input1)
			got, err := MarshalProto(tmp)
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

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
		want      ProtocBook
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
			want: ProtocBook{
				Books: []protoc.Book{
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
			want: ProtocBook{
				Books: []protoc.Book{
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
			want: ProtocBook{
				Books: []protoc.Book{
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
			want: ProtocBook{
				Books: []protoc.Book{
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
			assert.Equal(t, tC.want, got)
		})
	}
}

const (
	protoShit1 = "\n\x11978-5-389-21499-6\x12"
	protoShit2 = "#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x035\x9a\x99\x19@"
)

func TestUnmarshalProto(t *testing.T) {
	testCases := []struct {
		want      ProtocBook
		desc      string
		input1    *books.JSONBook
		input2    []ProtocBinary
		testError bool
	}{
		{
			desc:   "check valid",
			input1: &books.JSONBook{},
			want: ProtocBook{
				Books: []protoc.Book{
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
			input2: []ProtocBinary{
				{
					book: []byte(protoShit1 + protoShit2),
				}, {
					book: []byte(protoShit1 + protoShit2),
				},
			},
			testError: false,
		},
		{
			desc:   "check missing rate and size",
			input1: &books.JSONBook{},
			want: ProtocBook{
				Books: []protoc.Book{
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
			input2: []ProtocBinary{
				{
					book: []byte(protoShit1 + "\v ежика\x1a\x1dМюриель Барбери \xd9\x0f"),
				}, {
					book: []byte(protoShit1 + "\"Элегантность ежикa\x1a\x1dМюриель Барбери \xd9\x0f"),
				},
			},
			testError: false,
		},
		{
			desc:   "check one",
			input1: &books.JSONBook{},
			want: ProtocBook{
				Books: []protoc.Book{
					{
						ID:     "978-5-389-21499-6",
						Title:  " ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
					},
				},
			},
			input2: []ProtocBinary{
				{
					book: []byte("\n\x11978-5-389-21499-6\x12\v ежика\x1a\x1dМюриель Барбери \xd9\x0f"),
				},
			},
			testError: false,
		},
		{
			desc:   "check error",
			input1: &books.JSONBook{},
			want: ProtocBook{
				Books: []protoc.Book{
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
			input2: []ProtocBinary{
				{
					book: []byte("sfdsdfsdfsdf"),
				}, {
					book: []byte("34g4vt3g"),
				},
			},
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			protoBookNew, err := UnmarshalProto(tC.input2)
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				for i := range protoBookNew.Books {
					element := proto.Clone(&protoBookNew.Books[i]).(*protoc.Book)
					assert.Equal(t, tC.want.Books[i].GetID(), element.GetID())
					assert.Equal(t, tC.want.Books[i].GetAuthor(), element.GetAuthor())
					assert.Equal(t, tC.want.Books[i].GetRate(), element.GetRate())
					assert.Equal(t, tC.want.Books[i].GetSize(), element.GetSize())
					assert.Equal(t, tC.want.Books[i].GetYear(), element.GetYear())
					assert.Equal(t, tC.want.Books[i].GetTitle(), element.GetTitle())
				}
			}
		})
	}
}

func TestMarshalProto(t *testing.T) {
	testCases := []struct {
		input1    *books.JSONBook
		desc      string
		want      []ProtocBinary
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
			want: []ProtocBinary{
				{
					book: []byte(protoShit1 + protoShit2),
				},
				{
					book: []byte(protoShit1 + protoShit2),
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
					}, {
						ID:     "978-5-389-21499-6",
						Title:  "Элегантность ежика",
						Author: "Мюриель Барбери",
						Year:   2009,
						Size:   400,
					},
				},
			},
			want: []ProtocBinary{
				{
					book: []byte(protoShit1 + "#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x03"),
				},
				{
					book: []byte(protoShit1 + "#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x03"),
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
			want: []ProtocBinary{
				{
					book: []byte("\n\x11978-5-389-21499-6\x12\v ежика\x1a\x1dМюриель Барбери \xd9\x0f"),
				},
				{
					book: []byte("\n\x11978-5-389-21499-6\x12\"Элегантность ежикa\x1a\x1dМюриель Барбери \xd9\x0f"),
				},
			},
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tmp := ToProtobufStructure(tC.input1)
			got, err := tmp.MarshalProto()
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				for i, element := range got {
					assert.Equal(t, string(tC.want[i].book), string(element.book))
				}
			}
		})
	}
}

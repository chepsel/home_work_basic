package main

import (
	"testing"

	"github.com/chepsel/home_work_basic/hw09_serialize/books"
	"github.com/chepsel/home_work_basic/hw09_serialize/protoc"
	"github.com/stretchr/testify/assert"
)

func TestToProtobufStructure(t *testing.T) {
	testCases := []struct {
		want      *protoc.Book
		desc      string
		input1    *books.Book
		testError bool
	}{
		{
			desc: "check valid",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			testError: false,
		},
		{
			desc: "check missing rate",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			testError: false,
		},
		{
			desc: "check missing size and rate",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			want: &protoc.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
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

func TestMarshalProto(t *testing.T) {
	testCases := []struct {
		want   string
		desc   string
		input1 *books.Book
	}{
		{
			desc: "check valid",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			want: "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x035\x9a\x99\x19@",
		},
		{
			desc: "check missing rate",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			want: "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x03",
		},
		{
			desc: "check missing size and rate",
			input1: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			want: "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			temp := ToProtobufStructure(tC.input1)
			got, _ := MarshalProto(temp)
			assert.Equal(t, []byte(tC.want), got)
		})
	}
}

func TestUnmarshalProto(t *testing.T) {
	testCases := []struct {
		want      *books.Book
		desc      string
		input1    string
		testError bool
	}{
		{
			desc: "check valid",
			want: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
				Rate:   2.4,
			},
			input1:    "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x035\x9a\x99\x19@",
			testError: false,
		},
		{
			desc: "check missing rate",
			want: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
				Size:   400,
			},
			input1:    "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x03",
			testError: false,
		},
		{
			desc: "check missing size and rate",
			want: &books.Book{
				ID:     "978-5-389-21499-6",
				Title:  "Элегантность ежика",
				Author: "Мюриель Барбери",
				Year:   2009,
			},
			input1:    "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f",
			testError: false,
		},
		{
			desc:      "check wrong value",
			want:      &books.Book{},
			input1:    "\xd9\x0f(\x90\x03\n\x11978-5-389-21499-6\x12#Элегантность ежика\x12#Элегантность ежика4r3g",
			testError: true,
		},
		{
			desc:      "check non existing key",
			want:      &books.Book{},
			input1:    "32534f5fthb",
			testError: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ToProtobufStructure(&books.Book{})
			want := ToProtobufStructure(tC.want)
			err := UnmarshalProto([]byte(tC.input1), got)
			if tC.testError {
				if err == nil {
					t.Errorf("missing error")
				}
			} else {
				assert.Equal(t, want.GetID(), got.GetID())
				assert.Equal(t, want.GetRate(), got.GetRate())
				assert.Equal(t, want.GetAuthor(), got.GetAuthor())
				assert.Equal(t, want.GetSize(), got.GetSize())
				assert.Equal(t, want.GetTitle(), got.GetTitle())
				assert.Equal(t, want.GetYear(), got.GetYear())
			}
		})
	}
}

package books

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	const id, title, author string = "1", "2", "3"
	const year, size uint16 = 3, 2
	const rate float32 = 2.2
	want := &Book{id: id, title: title, author: author, year: year, size: size, rate: rate}
	got := NewBook(id, title, author, year, size, rate)
	assert.Equal(t, want, got)
}

func TestPrivateMakeLine(t *testing.T) { // tdt - шаблон готовый для использования(просто напиши tdt и жми enter)
	testCases := []struct {
		desc      string
		wantS     string
		wantU     uint16
		wantF     float32
		inputBook *Book
		inputS    string
		inputU    uint16
		inputF    float32
	}{
		{
			desc:      "Setters",
			inputBook: NewBook("21", "23", "23", 2, 2, 3.4),
			inputS:    "test",
			inputU:    2,
			inputF:    3.2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.wantS = tC.inputS
			tC.wantU = tC.inputU
			tC.wantF = tC.inputF
			tC.inputBook.SetID(tC.inputS)
			tC.inputBook.SetTitle(tC.inputS)
			tC.inputBook.SetAuthor(tC.inputS)
			tC.inputBook.SetYear(tC.inputU)
			tC.inputBook.SetSize(tC.inputU)
			tC.inputBook.SetRate(tC.inputF)
			gotid := tC.inputBook.ID()
			gottitle := tC.inputBook.Title()
			gotauthor := tC.inputBook.Author()
			gotyear := tC.inputBook.Year()
			gotsize := tC.inputBook.Size()
			gotrate := tC.inputBook.Rate()
			assert.Equal(t, tC.wantS, gotid)
			assert.Equal(t, tC.wantS, gottitle)
			assert.Equal(t, tC.wantS, gotauthor)
			assert.Equal(t, tC.inputU, gotsize)
			assert.Equal(t, tC.inputU, gotyear)
			assert.Equal(t, tC.inputF, gotrate)
		})
	}
}

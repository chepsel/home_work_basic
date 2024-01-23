package structure

import (
	"github.com/chepsel/home_work_basic/hw06_testing/structure/books"
)

type compareField int8

const (
	Year compareField = iota
	Size
	Rate
)

type Comparator struct {
	fieldCompare compareField
}

func (s compareField) String() string {
	switch s {
	case Year:
		return "year"
	case Size:
		return "size"
	case Rate:
		return "rate"
	default:
		return "unknown"
	}
}

func NewComparator(fieldComapre compareField) *Comparator {
	return &Comparator{fieldCompare: fieldComapre}
}

func (c Comparator) Compare(bookLeft, bookRight *books.Book) bool {
	switch {
	case c.fieldCompare == Year:
		return bookLeft.Year() > bookRight.Year()
	case c.fieldCompare == Size:
		return bookLeft.Size() > bookRight.Size()
	case c.fieldCompare == Rate:
		return bookLeft.Rate() > bookRight.Rate()
	default:
		return false
	}
}

func NewBook(i string, t string, a string, y uint16, s uint16, r float32) *books.Book {
	return books.NewBook(i, t, a, y, s, r)
}

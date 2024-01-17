package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw04_struct_comparator/books"
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
	}
	return "unknown"
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

func main() {
	year := NewComparator(Year)
	size := NewComparator(Size)
	rate := NewComparator(Rate)
	bookLeft := books.NewBook("978-5-389-21499-6", "Мюриель Барбери", "Элегантность ежика", 2009, 400, 2.4)
	bookRight := books.NewBook("172-8-335-00000-1", "Ибрагим Кимченымович", "Алкоголизм для чайников", 1969, 255, 3.9)

	fmt.Println("Left book", bookLeft)
	fmt.Println("Right book", bookRight)
	fmt.Println("compare", year.fieldCompare, "(left more then rigth):", year.Compare(bookLeft, bookRight))
	fmt.Println("compare", size.fieldCompare, "(left more then rigth):", size.Compare(bookLeft, bookRight))
	fmt.Println("compare", rate.fieldCompare, "(left more then rigth):", rate.Compare(bookLeft, bookRight))
}

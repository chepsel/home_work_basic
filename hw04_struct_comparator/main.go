package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw04_struct_comparator/internal"
)

type checkFormat string

const (
	Left  checkFormat = "left"
	Right checkFormat = "right"
)

type bookCompare struct {
	compareFormat checkFormat
	year          bool
	size          bool
	rate          bool
}

func intCompare(f checkFormat, l uint16, r uint16) bool {
	switch {
	case f == "left" && l > r:
		return true
	case f == "right" && r > l:
		return true
	default:
		return false
	}
}

func floatCompare(f checkFormat, l float32, r float32) bool {
	switch {
	case f == "left" && l > r:
		return true
	case f == "right" && r > l:
		return true
	default:
		return false
	}
}

func main() {
	enum := Right
	nextID := internal.IntSeq()
	bookLeft := internal.NewBook()
	bookRight := internal.NewBook()

	bookLeft.SetID(nextID())
	bookLeft.SetAuthor("Мюриель Барбери")
	bookLeft.SetTitle("Элегантность ежика")
	bookLeft.SetYear(2009)
	bookLeft.SetSize(400)
	bookLeft.SetRate(2.4)

	bookRight.SetID(nextID())
	bookRight.SetAuthor("Анохин, Сахаров")
	bookRight.SetTitle("Пособие тракториста")
	bookRight.SetYear(1969)
	bookRight.SetSize(255)
	bookRight.SetRate(3.9)

	compare := bookCompare{
		compareFormat: enum,
		year:          intCompare(enum, bookLeft.Year(), bookRight.Year()),
		size:          intCompare(enum, bookLeft.Size(), bookRight.Size()),
		rate:          floatCompare(enum, bookLeft.Rate(), bookRight.Rate()),
	}
	fmt.Println("compare type:", compare.compareFormat,
		"\nyear:", compare.year,
		"\nsize:", compare.size,
		"\nrate:", compare.rate,
		"\nbookLeft:", bookLeft,
		"\nbookRight:", bookRight)
}

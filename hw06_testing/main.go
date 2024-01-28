package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw06_testing/chessboard"
	"github.com/chepsel/home_work_basic/hw06_testing/fixapp"
	"github.com/chepsel/home_work_basic/hw06_testing/shapes"
	"github.com/chepsel/home_work_basic/hw06_testing/structure"
)

func main() {
	val, err := shapes.Circle.FigureArea(2)
	fmt.Println(val, err)

	year := structure.NewComparator(structure.Year)
	size := structure.NewComparator(structure.Size)
	rate := structure.NewComparator(structure.Rate)
	bookLeft := structure.NewBook("978-5-389-21499-6", "Мюриель Барбери", "Элегантность ежика", 2009, 400, 2.4)
	bookRight := structure.NewBook("172-8-335-00000-1", "Ибрагим Кимченымович", "Алкоголизм для чайников", 1969, 255, 3.9)

	fmt.Println("Left book", bookLeft)
	fmt.Println("Right book", bookRight)
	fmt.Println("compare year (left more then rigth):", year.Compare(bookLeft, bookRight))
	fmt.Println("compare size (left more then rigth):", size.Compare(bookLeft, bookRight))
	fmt.Println("compare rate (left more then rigth):", rate.Compare(bookLeft, bookRight))

	fmt.Println(chessboard.Chessboard(0))

	fixAppResult, _ := fixapp.FixApp("./fixapp/data.json")
	fmt.Println(fixAppResult)
}

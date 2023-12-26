package main

import (
	"fmt"
)

func makeLine(w int, uneven bool) (line string) {
	for i := 0; i < w; i++ {
		switch {
		case i%2 == 0 && uneven:
			line += "#"
		case uneven:
			line += " "
		case i%2 == 0 && !uneven:
			line += " "
		default:
			line += "#"
		}
	}
	line += "\n"
	return line
}

func checkBoard(s int) (result string) {
	h, w := s, s
	var line string
	var uneven bool
	for i := 0; i < h; i++ {
		switch {
		case i%2 == 0:
			uneven = false
		default:
			uneven = true
		}
		line = makeLine(w, uneven)
		result += line
	}
	return result
}

func getInput() (s int, err error) {
	var n int
	fmt.Println("Введите размер доски")
	n, err = fmt.Scanf("%d", &s)
	if n != 1 && err != nil {
		fmt.Printf("Ошибка '%v', ожидается циферное значение\nБудет использовано значение по умолчанию(8x8)\n", err)
		return s, err
	}
	return s, err
}

func main() {
	var size int
	var err error
	size, err = getInput()
	if err != nil {
		size = 8
	} else {
		fmt.Printf("Размер доски: %dx%d\n", size, size)
	}

	result := checkBoard(size)

	fmt.Println("------------------супердоска------------------")
	fmt.Println(result)
}

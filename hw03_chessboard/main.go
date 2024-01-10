package main

import (
	"fmt"
	"strings"
)

func makeLine(w int, firstNum int) string {
	var line strings.Builder
	for i := firstNum; i < w; i++ {
		switch {
		case i%2 == 0:
			line.WriteString("#")
		default:
			line.WriteString(" ")
		}
	}
	line.WriteString("\n")
	return line.String()
}

func checkBoard(s int) (result strings.Builder) {
	h, w := s, s
	for i := 0; i < h; i++ {
		switch {
		case i%2 == 0:
			result.WriteString(makeLine(w, 0))
		default:
			result.WriteString(makeLine(w+1, 1))
		}
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

	fmt.Println("------------------супердоска------------------")
	result := checkBoard(size)

	fmt.Println(result.String())
}

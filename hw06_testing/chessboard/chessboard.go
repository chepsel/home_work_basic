package chessboard

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

func Chessboard(size int) string {
	var internal int
	var err error
	if size <= 0 {
		internal, err = getInput()
		if err != nil {
			internal = 8
		}
	} else {
		internal = size
	}
	result := checkBoard(internal)
	return result.String()
}

package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrWrongValue = errors.New("wrong string, doesn't contain value")

func makeString(inputStr string) (map[string]int, error) {
	words := strings.Split(inputStr, " ")
	var list []string
	for _, word := range words {
		value := strings.ReplaceAll(validateString(word), "\x00", "")
		if len(value) > 0 {
			list = append(list, value)
		}
	}
	if len(list) == 0 {
		return nil, ErrWrongValue
	}
	duplicatesCount := make(map[string]int)
	for _, item := range list {
		duplicatesCount[item]++
	}
	return duplicatesCount, nil
}

func validateString(value string) string {
	clean := func(r rune) rune {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				return unicode.ToLower(r)
			}
			return r
		}
		return rune(0)
	}
	return strings.Map(clean, value)
}

func main() {
	msg := "один два: три,    четыре пять шесть, один два. кукареку \n    Четыре восемь игнат димас шесть"
	words, err := makeString(msg) // делаем массив из текста
	if err != nil {
		fmt.Println(err)
	} else {
		for k, v := range words {
			fmt.Printf("%s:%d\n", k, v)
		}
	}
}

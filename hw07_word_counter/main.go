package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	ErrWrongValue = errors.New("wrong string, doesn't contain value")
)

func countWords(words string) (map[string]int, error) {
	list, err := makeString(words)
	if err != nil {
		return nil, err
	}
	duplicatesCount := make(map[string]int)
	for _, item := range list {
		_, exist := duplicatesCount[item]
		if exist {
			duplicatesCount[item] += 1
		} else {
			duplicatesCount[item] = 1
		}

	}
	return duplicatesCount, nil
}

func validateString(value string) string {
	clean := func(r rune) rune {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				return unicode.ToLower(r)
			} else {
				return r
			}
		}
		return rune(0)
	}
	return strings.Map(clean, value)
}

func makeString(input string) ([]string, error) {
	words := strings.Split(input, " ")
	var result []string
	for _, word := range words {
		value := strings.Replace(validateString(word), "\x00", "", -1)
		switch {
		case len(value) > 0:
			result = append(result, value)
		}

	}
	if len(result) > 0 {
		return result, nil
	} else {
		return nil, ErrWrongValue
	}
}

func printResults(duplicatesMap map[string]int) {
	for k, v := range duplicatesMap {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func main() {
	msg := "один два: три,    четыре пять шесть, один два. кукареку \n    Четыре восемь игнат димас шесть"
	words, err := countWords(msg) // делаем массив из текста
	if err != nil {
		fmt.Println(err)
	} else {
		printResults(words)
	}
}

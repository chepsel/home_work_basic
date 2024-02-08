package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

func NewCryptoRand(limit int) int {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	if err != nil {
		panic(err)
	}
	return int(safeNum.Int64())
}

func randArr(array *[]int, elements int) []int {
	for i := 0; i < elements; i++ {
		(*array)[i] = NewCryptoRand(elements * 2)
	}
	sort.Ints(*array)
	return (*array)
}

func binarySearch(array *[]int, lookingFor int) (int, bool) {
	var midle, currentValue, o int

	min := 0
	top := len(*array) - 1

	for min <= top {
		o++
		midle = (min + top) / 2
		currentValue = (*array)[midle]
		if currentValue == lookingFor {
			fmt.Printf("O(%d)\n", o)
			return currentValue, true
		}
		if currentValue > lookingFor {
			top = midle - 1
		} else {
			min = midle + 1
		}
	}
	fmt.Printf("O(%d)\n", o)
	return 0, false
}

func main() {
	const elements int = 50
	lookingFor := NewCryptoRand(elements)
	array := make([]int, elements)

	randArr(&array, elements)

	assumption, result := binarySearch(&array, lookingFor)
	if result {
		fmt.Println("Found: " + strconv.Itoa(assumption))
		fmt.Println(array)
	} else {
		fmt.Println("Can't found:" + strconv.Itoa(lookingFor))
		fmt.Println(array)
	}
}

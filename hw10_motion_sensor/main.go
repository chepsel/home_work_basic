package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	chanelStat := make(chan uint64)
	ageregatedStat := make(chan uint64)
	go collectStat(chanelStat, 60)
	go aggregateStat(chanelStat, ageregatedStat)
	for task := range ageregatedStat {
		fmt.Println("AverageAllocMemory:", task, " Bytes")
	}
}

func CryptoRand(limit int) uint64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(int64(limit)))
	if err != nil {
		return 0
	}
	return uint64(safeNum.Int64())
}

func collectStat(chanelStat chan<- uint64, seconds int) {
	defer close(chanelStat)
	to := time.After(time.Duration(seconds) * time.Second)
	for {
		select {
		case <-to:
			return
		default:
			chanelStat <- CryptoRand(9999999)
		}
	}
}

func aggregateStat(chanelStat <-chan uint64, ageregatedStat chan<- uint64) {
	defer close(ageregatedStat)
	var sum uint64
	var count uint64
	for task := range chanelStat {
		count++
		sum += task
		if count == 10 {
			result := sum / count
			ageregatedStat <- result
			sum = 0
			count = 0
		}
	}
}

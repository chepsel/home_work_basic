package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	chanelStat := make(chan uint64, 1)
	ageregatedStat := make(chan uint64, 1)
	endSignal := make(chan bool, 1)
	go collectStat(chanelStat, endSignal)
	go aggregateStat(chanelStat, ageregatedStat)
	go stopper(60, endSignal)
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

func stopper(seconds int, endSignal chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	endSignal <- true
}

func collectStat(chanelStat chan<- uint64, endSignal <-chan bool) {
	defer close(chanelStat)
	chanelStat <- CryptoRand(9999999)
	for {
		select {
		case <-endSignal:
			return
		case <-time.After(time.Second):
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

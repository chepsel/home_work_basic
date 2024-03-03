package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	chanelStat := make(chan uint64)
	ageregatedStat := make(chan uint64)
	loopCounter := 60
	go collectStat(chanelStat, loopCounter)
	go aggregateStat(chanelStat, ageregatedStat)
	go printStat(ageregatedStat)

	time.Sleep(time.Second * time.Duration(loopCounter+1))
}

func collectStat(chanelStat chan<- uint64, loopCounter int) {
	for i := 0; i < loopCounter; i++ {
		chanelStat <- PrintMemUsage()
		time.Sleep(time.Second)
	}
	close(chanelStat)
}

func aggregateStat(chanelStat <-chan uint64, ageregatedStat chan<- uint64) {
	var sum uint64
	var count uint64
	for {
		task, ok := <-chanelStat
		if ok {
			count++
			sum += task
			if count == 10 {
				result := sum / count
				ageregatedStat <- result
				sum = 0
				count = 0
			}
		} else {
			close(ageregatedStat)
			break
		}
	}
}

func printStat(ageregatedStat <-chan uint64) {
	for {
		task, ok := <-ageregatedStat
		if ok {
			fmt.Println("AverageAllocMemory:", task, " Bytes")
		} else {
			break
		}
	}
}

func PrintMemUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.TotalAlloc
}

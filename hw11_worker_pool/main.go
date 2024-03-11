package main

import (
	"fmt"
	"sync"
)

func main() {
	var tasksCounter int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	gouroutineNumber := 4
	gourouteQueue := make(chan int, gouroutineNumber)
	for i := 0; i < gouroutineNumber; i++ {
		SomeTask(gourouteQueue, &wg, &mu, i, &tasksCounter)
	}
	wg.Wait()
	for i := 0; i < cap(gourouteQueue); i++ {
		fmt.Println("Function (", <-gourouteQueue, ") ended")
	}
	close(gourouteQueue)
	fmt.Println("Result is:", tasksCounter)
}

func SomeTask(gourouteQueue chan<- int, wg *sync.WaitGroup, mu *sync.Mutex, funcNum int, tasksCounter *int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Function", funcNum, "started")
		for i := 0; i < 5000; i++ {
			mu.Lock()
			*tasksCounter++
			mu.Unlock()
		}
		gourouteQueue <- funcNum
	}()
}

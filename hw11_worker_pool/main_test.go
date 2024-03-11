package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeTask(t *testing.T) {
	testCases := []struct {
		want   int
		desc   string
		input1 int
	}{
		{
			desc:   "check gorutine - 5",
			input1: 5,
			want:   25000,
		},
		{
			desc:   "check gorutine - 4",
			input1: 4,
			want:   20000,
		},
		{
			desc:   "check gorutine - 1",
			input1: 1,
			want:   5000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var tasksCounter int
			wg := sync.WaitGroup{}
			mu := sync.Mutex{}
			gourouteQueue := make(chan int, tC.input1)
			for i := 0; i < tC.input1; i++ {
				SomeTask(gourouteQueue, &wg, &mu, i, &tasksCounter)
			}
			wg.Wait()
			close(gourouteQueue)
			assert.Equal(t, tasksCounter, tC.want)
		})
	}
}

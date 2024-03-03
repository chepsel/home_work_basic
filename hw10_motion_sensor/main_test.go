package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCollectStat(t *testing.T) {
	testCases := []struct {
		want      bool
		desc      string
		input1    int
		testError bool
	}{
		{
			desc:      "check gorutine - ok",
			input1:    5,
			want:      true,
			testError: false,
		},
		{
			desc:      "check gorutine - error",
			input1:    5,
			want:      false,
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chanelStat := make(chan uint64)
			go collectStat(chanelStat, tC.input1)
			if tC.testError {
				for i := 0; i < tC.input1; i++ {
					_, ok := <-chanelStat
					if !ok {
						break
					}
				}
				_, ok := <-chanelStat
				assert.Equal(t, tC.want, ok)
			} else {
				for i := 0; i < tC.input1; i++ {
					_, ok := <-chanelStat
					assert.Equal(t, tC.want, ok)
				}
			}
		})
	}
}

func TestAggregateStat(t *testing.T) {
	testCases := []struct {
		want      uint64
		desc      string
		input1    []uint64
		testError bool
	}{
		{
			desc:      "check gorutine - ok",
			input1:    []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      5,
			testError: false,
		},
		{
			desc:      "check gorutine - ok 2",
			input1:    []uint64{12, 4, 3, 40, 5, 6, 1, 8, 0, 20},
			want:      9,
			testError: false,
		},
		{
			desc:      "check gorutine - big array",
			input1:    []uint64{12, 4, 3, 40, 5, 6, 1, 8, 100, 20, 12, 4, 3, 40, 5, 6, 1, 8, 100, 20},
			want:      19,
			testError: false,
		},
		{
			desc:      "check gorutine - error",
			input1:    []uint64{12},
			want:      19,
			testError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chanelStat := make(chan uint64)
			ageregatedStat := make(chan uint64)
			go aggregateStat(chanelStat, ageregatedStat)
			go func() {
				for _, v := range tC.input1 {
					chanelStat <- v
				}
				close(chanelStat)
			}()
			go func() {
				for {
					got, ok := <-ageregatedStat
					if ok {
						assert.Equal(t, tC.want, got)
					} else {
						if tC.testError {
							assert.Equal(t, ok, false)
						}
						break
					}
				}
			}()
			time.Sleep(time.Second)
		})
	}
}

func TestPrintMemUsage(t *testing.T) {
	got := PrintMemUsage()
	if got <= 0 {
		t.Errorf("error")
	}
}

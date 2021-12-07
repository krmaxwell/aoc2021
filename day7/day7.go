package day7

import (
	"math"
)

func MinimumFuel(positions []int) int {
	minfuel := math.MaxInt
	maxPosition := maxInt(positions)
	fuel := make([]int, maxPosition+1)
	for pos := 0; pos <= maxPosition; pos++ {
		for _, crab := range positions {
			if crab > pos {
				fuel[pos] += crab - pos
			} else {
				fuel[pos] += pos - crab
			}
		}
		if fuel[pos] < minfuel {
			minfuel = fuel[pos]
		}
	}
	return minfuel
}

func maxInt(data []int) int {
	max := 0
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

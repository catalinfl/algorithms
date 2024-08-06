package problems

import (
	"slices"
)

func GreedyMinimumPushes(word string) int {
	freq := make([]int, 26)

	for _, c := range word {
		freq[int(c-'a')]++
	}

	slices.Sort(freq)

	res := 0
	multiplier := 1

	counter := 8

	for i := 25; i >= 0; i-- {
		res += freq[i] * multiplier

		counter--

		if counter == 0 {
			multiplier++
			counter = 8
		}

	}

	return res

}

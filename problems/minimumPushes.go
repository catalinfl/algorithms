package problems

import (
	"sort"
)

func MinimumPushes(word string) int {

	st := make(map[byte]int)

	mapped := make(map[byte]int)

	for i := 0; i < len(word); i++ {
		st[word[i]]++
	}

	type kv struct {
		Key   byte
		Value int
	}

	var sortedSlice []kv

	for k, v := range st {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value > sortedSlice[j].Value
	})

	tastaValue := 1
	count := 0

	for i := 0; i < len(sortedSlice); i++ {

		count++

		mapped[sortedSlice[i].Key] = tastaValue

		if count > 7 {
			tastaValue++
			count = 0
		}

	}

	totalPushes := 0

	for i := 0; i < len(word); i++ {

		if word[i] == '1' || word[i] == '*' || word[i] == '0' || word[i] == '#' {
			totalPushes += 1
			continue
		}

		totalPushes += mapped[word[i]]

	}

	return totalPushes

}

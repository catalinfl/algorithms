package problems

import (
	"fmt"
	"sort"
	"strconv"
)

func FindRelativeRanks(score []int) []string {

	copy_score := make([]int, len(score))

	copy(copy_score, score)

	sort.Slice(copy_score, func(i, j int) bool {
		return copy_score[i] > copy_score[j]
	})

	medal_map := make(map[int]string)

	for rank, val := range copy_score {
		var metal string

		switch rank {
		case 0:
			metal = "Gold Medal"
		case 1:
			metal = "Silver Medal"
		case 2:
			metal = "Bronze Medal"
		default:
			metal = strconv.Itoa(rank + 1)
		}

		medal_map[val] = metal
	}

	result := []string{}

	for _, points := range score {
		result = append(result, medal_map[points])
	}

	fmt.Println(result)

	return result

}

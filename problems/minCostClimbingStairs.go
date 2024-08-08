package problems

func MinCostClimbingStairs(cost []int) int {

	len := len(cost)

	for i := 2; i < len; i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}

	return min(cost[len-1], cost[len-2])
}

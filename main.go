package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))

}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var backTracking func([]int, int)

	backTracking = func(combination []int, nextIndex int) {
		total := sum(combination...)

		if target <= total {
			if target == total {
				copyCombination := make([]int, len(combination))
				copy(copyCombination, combination)
				result = append(result, copyCombination)
			}

			return
		}

		for i := nextIndex; i < len(candidates); i++ {
			backTracking(append(combination, candidates[i]), i)
		}

	}

	backTracking(make([]int, 0), 0)

	return result

}

func sum(nums ...int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	return total
}

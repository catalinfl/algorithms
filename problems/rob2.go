package problems

import "fmt"

func Rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	var rob func(int, []int, map[int]int) int

	rob = func(pos int, arr []int, memo map[int]int) int {

		if v, ok := memo[pos]; ok {
			return v
		}

		if pos >= len(arr) {
			return 0
		}

		a := rob(pos+1, arr, memo)
		b := rob(pos+2, arr, memo) + arr[pos]
		ans := max(a, b)
		memo[pos] = ans
		fmt.Println(ans)
		return ans
	}

	memo1 := make(map[int]int)

	memo2 := make(map[int]int)

	return max(rob(1, nums, memo1), rob(0, nums[:len(nums)-1], memo2))

}

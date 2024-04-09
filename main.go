package main

import (
	"fmt"
	"sort"
)

func main() {

	var prices []int = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	l := len(prices)

	if l == 0 {
		return 0
	}

	b := make([]int, len(prices))

	copy(b, prices)

	sort.Ints(b)

	minPrice := b[0]

	for i := 0; i < l; i++ {
		b[i] = 0
		fmt.Println(b[i])
	}

	for i := 0; i < l; i++ {
		b[i] = prices[i] - minPrice
	}

	sort.Ints(b)

	return b[len(b)-1]

}

package dynamicprogramming

func Sum(n int) int {
	a := make([]int, n+1)

	a[0] = 0

	for i := 1; i <= n; i++ {
		a[i] = a[i-1] + i
	}

	return a[n]
}

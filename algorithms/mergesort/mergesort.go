package mergesort

import "fmt"

func Merge(arr []int, l int, m int, r int) {
	fmt.Printf("Primul merge, left:%d, middle: %d, right:%d \n", l, m, r)

	var i, j, k int
	var n1 int = m - l + 1
	var n2 int = r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i = 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for j = 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	i = 0
	j = 0
	k = l

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[i]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func MergeSort(arr []int, l int, r int) {
	fmt.Printf("Mergesort left: %d, right: %d \n", l, r)

	if l < r {
		m := l + (r-l)/2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		Merge(arr, l, m, r)
	}

}

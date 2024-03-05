package main

import (
	"fmt"

	"github.com/catalinfl/algorithms/algorithms/mergesort"
)

func main() {
	var arr []int = []int{13, 9, 4, 12}

	mergesort.MergeSort(arr, 0, len(arr)-1)

	for _, ar := range arr {
		fmt.Println(ar)
	}

}

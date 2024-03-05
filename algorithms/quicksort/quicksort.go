package quicksort

// func partition(arr []int, low int, high int) int {

// 	fmt.Printf("Partition: low: %d, high: %d \n", low, high)

// 	var pivot int = arr[high]
// 	var i int = low - 1

// 	for j := low; j <= high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	fmt.Printf("This partition returns %d \n", i+1)
// 	fmt.Printf("For this partition we have next order of numbers: \n")
// 	for _, ar := range arr {
// 		fmt.Printf("%d, ", ar)
// 	}
// 	fmt.Printf("\n")
// 	return i + 1
// }

// func QuickSort(arr []int, low int, high int) {

// 	fmt.Printf("Quicksort low: %d, high: %d \n", low, high)

// 	if low < high {
// 		var pi int = partition(arr, low, high)

// 		QuickSort(arr, low, pi-1)
// 		QuickSort(arr, pi+1, high)

// 	}
// }

func partition(arr []int, low int, high int) int {

	var pivot int = arr[high]
	var i = low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		var pi int = partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

package problems

func kthDistinct(arr []string, k int) string {
	res, strFrequency := "", make(map[string]int)

	for i := 0; i < len(arr); i++ {
		strFrequency[arr[i]]++
	}

	for i := 0; i < len(arr); i++ {
		if strFrequency[arr[i]] == 1 {
			k--
		}

		if k == 0 {
			res = arr[i]
			break
		}

	}

	return res

}

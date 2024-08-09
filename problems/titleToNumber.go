package problems

func TitleToNumber(columnTitle string) int {

	// freq := make([]int, 26)

	sum := 0
	for i := 0; i < len(columnTitle); i++ {
		sum = (int(columnTitle[i]) - 64) + sum*26
	}

	return sum
}

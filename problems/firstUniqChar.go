package problems

func firstUniqChar(s string) int {
	st := make(map[string]int, len(s))

	for i := 0; i < len(s); i++ {
		st[string(s[i])]++
	}

	for i := 0; i < len(s); i++ {
		if st[string(s[i])] == 1 {
			return i
		}
	}

	return -1

}

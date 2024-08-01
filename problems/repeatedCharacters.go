package problems

func repeatedCharacter(s string) byte {
	st := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		st[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if st[s[i]] > 1 {
			return s[i]
		}
	}

	return 'z'

}

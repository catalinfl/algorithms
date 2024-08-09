package problems

func ConvertToTitle(columnNumber int) string {

	str := ""

	for columnNumber != 0 {

		str += string(byte('A' + (columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}

	banan := reverse(str)

	return banan
}

func reverse(str string) string {

	bytes := []byte(str)
	i := 0
	j := len(str) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)

}

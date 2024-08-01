package problems

import (
	"strconv"
)

func CountSeniors(details []string) int {
	count := 0

	for _, detail := range details {
		age := detail[11:13]

		ageInt, _ := strconv.Atoi(age)

		if ageInt > 60 {
			count++
		}

	}

	return count

}

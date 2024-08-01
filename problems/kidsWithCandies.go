package problems

func KidsWithCandies(candies []int, extraCandies int) []bool {

	candiesMax := 0

	var candiesBool []bool

	for i := 0; i < len(candies); i++ {
		if candiesMax < candies[i] {
			candiesMax = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		candyBool := false
		if candies[i]+extraCandies >= candiesMax {
			candyBool = true
		}
		candiesBool = append(candiesBool, candyBool)
	}

	return candiesBool

}

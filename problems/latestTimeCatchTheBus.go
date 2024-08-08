package problems

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	firstBus := 0
	secondBus := 1

	var capacity_temp int = capacity

	for secondBus < len(buses) {
		for i := 0; i < len(passengers); i++ {

			if passengers[i] < buses[firstBus] {
				if capacity_temp == 1 {
					return passengers[i] - 1
				}
				capacity_temp--
				continue
			}

			if buses[firstBus] < passengers[i] && buses[secondBus] > passengers[i] {
				if capacity_temp == 1 {
					return passengers[i] - 1
				}
				capacity_temp--
				continue
			} else {
				firstBus++
				secondBus++
				i = i - 1
			}
		}
	}

	return passengers[0] - 1

}

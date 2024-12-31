package main

func predictPartyVictory(senate string) string {
	radiantQueue := []int{}
	direQueue := []int{}
	n := len(senate)

	//enqueue to both queues
	for index, ele := range senate {
		if ele == 'R' {
			radiantQueue = append(radiantQueue, index)
		} else {
			direQueue = append(direQueue, index)
		}
	}

	//simulate voting
	for len(radiantQueue) > 0 && len(direQueue) > 0 {
		radIndex := radiantQueue[0]
		direIndex := direQueue[0]

		radiantQueue = radiantQueue[1:]
		direQueue = direQueue[1:]

		if radIndex < direIndex {
			radiantQueue = append(radiantQueue, radIndex+n)
		} else {
			direQueue = append(direQueue, direIndex+n)
		}
	}

	if len(radiantQueue) > 0 {
		return "Radiant"
	}
	return "Dire"

}

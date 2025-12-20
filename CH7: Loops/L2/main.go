package main

func maxMessages(thresh int) int {
	var msgCost int

	for i := 0; ; i++ {
		msgCost += 100 + i

		if msgCost > thresh {
			return i
		}
	}

}

package main

func bulkSend(numMessages int) float64 {
	var msgCost float64
	for i := 0; i < numMessages; i++ {

		msgCost += 1.0 + float64(i)*0.01

	}
	return msgCost
}

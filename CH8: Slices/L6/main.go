package main

func getMessageCosts(messages []string) []float64 {
	// ?

	messageCost := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		messageCost[i] = float64(len(messages[i])) * 0.01
	}

	return messageCost
}

package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {

	var stringArray [3]string
	var costArray [3]int

	stringArray[0] = primary
	stringArray[1] = secondary
	stringArray[2] = tertiary

	costArray[0] = len(primary)
	for i := 1; i < len(costArray); i++ {
		costArray[i] = costArray[i-1] + len(stringArray[i])
	}

	return stringArray, costArray
}

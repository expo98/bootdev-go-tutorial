package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	var totalCost int

	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		fmt.Println("Couldn't send message: ", err)
		return 0, err
	}
	totalCost += cost

	cost, err = sendSMS(msgToSpouse)
	if err != nil {
		fmt.Println("Couldn't send message:", err)
		return 0, err
	}
	totalCost += cost

	return totalCost, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

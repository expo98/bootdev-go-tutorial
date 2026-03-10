package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {

	if plan == planPro {

		messagesSlice := messages[:]
		return messagesSlice, nil

	}

	if plan == planFree {

		messagesSlice := messages[:2]
		return messagesSlice, nil
	}

	var err error = errors.New("unsupported plan")
	return nil, err
}

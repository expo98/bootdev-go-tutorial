package main

import "fmt"

func (e email) cost() int {
	body := e.body
	var bodyCost int

	if !e.isSubscribed {
		bodyCost = 5
	} else {
		bodyCost = 2
	}

	return len(body) * bodyCost
}

func (e email) format() string {
	var secondPart string

	if !e.isSubscribed {
		secondPart = "Not Subscribed"
	} else {
		secondPart = "Subscribed"
	}

	formattedMessage := fmt.Sprintf("'%s' | %s", e.body, secondPart)

	return formattedMessage
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

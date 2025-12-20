package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

// ?

func (d directMessage) importance() int {
	score := d.priorityLevel

	if d.isUrgent {
		score = 50
	}

	return score
}

func (g groupMessage) importance() int {
	score := g.priorityLevel

	return score
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch noti := n.(type) {
	case (directMessage):
		return noti.senderUsername, noti.importance()
	case (groupMessage):
		return noti.groupName, noti.importance()
	case (systemAlert):
		return noti.alertCode, n.importance()
	default:
		return "", 0
	}
}

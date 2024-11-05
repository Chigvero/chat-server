package entities

import "time"

type SendMessage struct {
	ID           int
	ChatId       int
	FromUserName string
	MessageText  string
	Timestamp    time.Time
}

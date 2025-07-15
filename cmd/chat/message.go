package chat

import (
	"strconv"
	"time"

	"www.github.com/M1ralai/tcp-server/cmd/users"
)

type Message struct {
	Time    string
	Sender  string
	Message string
}

func NewMessage(sender users.User, message string) Message {
	hour, minute, second := time.Now().Clock()

	return Message{
		Time:    strconv.Itoa(hour) + strconv.Itoa(minute) + strconv.Itoa(second),
		Sender:  sender.Username,
		Message: message,
	}
}

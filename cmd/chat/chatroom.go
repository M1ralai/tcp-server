package chat

import "www.github.com/M1ralai/tcp-server/cmd/users"

type Chatroom struct {
	Name        string
	Message     chan (Message)
	Connections []users.User
	admin       *users.User
}

func NewChatroom(name string, creator *users.User) *Chatroom {
	return &Chatroom{
		Name:        name,
		Message:     make(chan Message),
		Connections: nil,
		admin:       creator,
	}
}

func (cr *Chatroom) SendMessage() {
	for {
		msg := <-cr.Message
		for i := range cr.Connections {
			if msg.Sender != cr.Connections[i].Username {
				cr.Connections[i].Conn.Write([]byte(msg.Message))
			}
		}
	}
}

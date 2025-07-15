package server

import (
	"errors"
	"net"
	"unicode"

	"www.github.com/M1ralai/tcp-server/cmd/chat"
	"www.github.com/M1ralai/tcp-server/cmd/users"
)

func (s *Server) loginUser(unm string, psw string, conn net.Conn) (*users.User, error) {
	for i := range s.users {
		if s.users[i].Username == unm && s.users[i].Password == psw && !s.users[i].LoggedIn {
			s.users[i].Conn = conn
			s.users[i].LoggedIn = true
			s.users[i].ChatRoom = "default"
			return &s.users[i], nil
		}
	}
	return nil, IUOP
}

func (s *Server) newChatRoom(cnm string, admin users.User) (*chat.Chatroom, error) {
	for i := range cnm {
		if !unicode.IsLetter(rune(cnm[i])) {
			return nil, errors.New("just use letters when giving a name")
		}
	}
	for i := range s.Chatrooms {
		if s.Chatrooms[i].Name == cnm {
			return nil, errors.New("there is a chatroom already using this name")
		}
	}
	return chat.NewChatroom(cnm, &admin), nil
}

package server

import (
	"net"

	"www.github.com/M1ralai/tcp-server/cmd/users"
)

func (s *Server) loginUser(unm string, psw string, conn net.Conn) (*users.User, error) {
	for i := range s.users {
		if s.users[i].Username == unm && s.users[i].Password == psw && !s.users[i].LoggedIn {
			s.users[i].Conn = conn
			s.users[i].LoggedIn = true
			return &s.users[i], nil
		}
	}
	return nil, IUOP
}

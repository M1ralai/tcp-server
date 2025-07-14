package server

import "net"

func (s *Server) welcomePage(conn net.Conn) {
	conn.Write([]byte("type :login for login already exist account or :register for create a new user \n"))
	str := read(conn)
	switch str {
	case ":login":
		conn.Write([]byte("You are redirecting to a login page\n"))
	case ":register":
		conn.Write([]byte("You are redirecting to a register page\n"))
	default:
		conn.Write([]byte("Please choose one of the options\n"))
	}
}

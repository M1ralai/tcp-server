package server

import (
	"log"
	"net"
)

const ()

type Server struct {
	serverAddr string
}

func NewServer(addr string) *Server {
	return &Server{
		serverAddr: addr,
	}
}

func (s *Server) Run() {
	ls, err := net.Listen("tcp", s.serverAddr)
	if err != nil {
		log.Println("error occured when creating listener for tcp server: ", err)
		return
	}
	s.acceptLoop(ls)
}

func (s *Server) acceptLoop(ls net.Listener) {
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Println("error occured when pair trying to connect server: ", err)
		}
		conn.Write([]byte("Hello Meine Freunde \n"))
		conn.Close()
	}
}

package server

import (
	"log"
	"net"

	"www.github.com/M1ralai/tcp-server/cmd/client"
)

const ()

type Server struct {
	serverAddr string
	users      []client.Client
}

func NewServer(addr string) *Server {
	return &Server{
		serverAddr: addr,
		users:      []client.Client{},
	}
}

func (s *Server) Run() {
	ls, err := net.Listen("tcp", s.serverAddr)
	if err != nil {
		log.Println("error occured when creating listener for tcp server: ", err)
		return
	} else {
		log.Println("tcp server started at ", s.serverAddr, "")
	}
	s.acceptLoop(ls)
}

func (s *Server) acceptLoop(ls net.Listener) {
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Println("error occured when pair trying to connect server: ", err)
		} else {
			log.Println("connection established from: ", conn.RemoteAddr())
		}
		conn.Write([]byte("Wrlcome to the M1ralai's tcp chat server \n"))
		s.welcomePage(conn)
	}
}

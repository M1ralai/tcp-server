package server

import (
	"log"
	"net"
)

func read(conn net.Conn) string {
	b := make([]byte, 32)
	n, err := conn.Read(b)
	str := string(b[:n-2])
	if err != nil {
		log.Println("error occured when reading message from: ", conn.RemoteAddr())
	} else {
		log.Println(conn.LocalAddr(), " write ", str, " to the chat")
	}
	return str
}

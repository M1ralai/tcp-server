package server

import (
	"errors"
	"log"
	"net"
)

var IUOP = errors.New("invalid username or password")
var ConnectionLost = errors.New("connection is lost because of client side problem")

func read(conn net.Conn, function string) (string, error) {
	b := make([]byte, 32)
	n, err := conn.Read(b)
	if n == 0 {
		return " ", ConnectionLost
	}
	str := string(b[:n-2])
	if err != nil {
		log.Println("error occured when reading message from: ", conn.RemoteAddr())
	} else {
		log.Println(conn.LocalAddr(), " write ", str, " to the ", function)
	}
	return str, nil
}

package main

import "www.github.com/M1ralai/tcp-server/cmd/server"

func main() {
	s := server.NewServer(":6969")
	s.Run()
}

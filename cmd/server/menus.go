package server

import (
	"log"
	"net"
	"unicode"

	"www.github.com/M1ralai/tcp-server/cmd/users"
)

func (s *Server) welcomePage(conn net.Conn) {
	for {
		conn.Write([]byte("type :login for login already exist account or :register for create a new user \n"))
		str, err := read(conn, "welcomePage")
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		switch str {
		case ":login":
			conn.Write([]byte("You are redirecting to a login page\n"))
			s.loginPage(conn)
			return
		case ":register":
			conn.Write([]byte("You are redirecting to a register page\n"))
			s.registerPage(conn)
			return
		default:
			conn.Write([]byte("Please choose one of the options\n"))
		}
	}
}

func (s *Server) registerPage(conn net.Conn) {
	for {
		conn.Write([]byte("give me your username"))
		unm, err := read(conn, "regsiterPage")
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		conn.Write([]byte("give me your password"))
		psw, err := read(conn, "regsiterPage")
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		if len(unm) > 4 || len(psw) > 4 {
			for i := range unm {
				if !unicode.IsLetter(rune(unm[i])) {
					conn.Write([]byte("username must be all letter"))
					break
				} else {
					u := users.NewUser(unm, psw)
					s.users = append(s.users, *u)
					return
				}
			}
		} else {
			conn.Write([]byte("username or password cannot be shorter than 4 char"))
		}
	}
}

func (s *Server) loginPage(conn net.Conn) {
	for {
		conn.Write([]byte("give me your username"))
		unm, err := read(conn, "loginPage")
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		conn.Write([]byte("give me your password"))
		psw, err := read(conn, "loginPage")
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		u, err := s.loginUser(unm, psw, conn)
		if err != nil {
			conn.Write([]byte("Username or Password is invalid"))
			log.Println(conn.RemoteAddr(), " tried to login but ", err)
		} else {
			s.chat(u)
			return
		}
	}
}

func (s *Server) chat(u *users.User) {
	for {
		str, err := read(u.Conn, "chat")
		if err != nil {
			log.Println(err)
			u.LogOut()
			return
		}

	}
}

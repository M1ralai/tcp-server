package users

import "net"

type User struct {
	Username string
	Password string
	LoggedIn bool
	ChatRoom string
	Conn     net.Conn
}

func NewUser(name string, password string) *User {
	return &User{
		Username: name,
		Password: password,
		ChatRoom: "default",
		LoggedIn: false,
		Conn:     nil,
	}
}

func (u *User) LogOut() {
	u.Conn.Close()
	u.LoggedIn = false
	u.Conn = nil
}

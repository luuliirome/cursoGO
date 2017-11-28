package domain

import (
	"fmt"
)

type User struct {
	Name       string
	Mail       string
	Nickname   string
	Contraseña string
	Tweets     []*Tweet
	Followers  []*User
	Following  []*User
}

func NewUser(name string, mail string, nick string, contraseña string) (*User, error) {
	if name == "" || mail == "" || contraseña == "" || nick == "" {
		var err error = fmt.Errorf("invalid user")
		return nil, err
	}
	return &User{Name: name, Mail: mail, Contraseña: contraseña, Nickname: nick}, nil
}

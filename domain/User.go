package domain

import (
	"fmt"
)

type User struct {
	Name             string
	CantidadDeTweets int
}

func NewUser(name string) (*User, error) {
	if name == "" {
		var err error = fmt.Errorf("invalid username")
		return nil, err
	}
	return &User{Name: name}, nil
}

package domain

import (
	"fmt"
	"time"
)

var currentId int = 0

// Tweet struct es la estructura tweet
type Tweet struct {
	Id   int
	User *User
	Text string
	Date *time.Time
}

// NewTweet crea un nuevo tweet
func NewTweet(user *User, text string) (Tweet, error) {
	var timeNow = time.Now()

	if text == "" {
		var err error = fmt.Errorf("text is required")
		return Tweet{}, err
	}

	if user == nil {
		var err error = fmt.Errorf("user is required")
		return Tweet{}, err
	}

	currentId++
	return Tweet{User: user, Text: text, Date: &timeNow, Id: currentId}, nil
}

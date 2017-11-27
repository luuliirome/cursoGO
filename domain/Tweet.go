package domain

import (
	"fmt"
	"time"
)

// Tweet struct es la estructura tweet
type Tweet struct {
	User *User
	Text string
	Date *time.Time
}

// NewTweet crea un nuevo tweet
func NewTweet(user *User, text string) (*Tweet, error) {
	var timeNow = time.Now()

	if text == "" {
		var err error = fmt.Errorf("text is required")
		return nil, err
	}

	if user == nil {
		var err error = fmt.Errorf("user is required")
		return nil, err
	}

	return &Tweet{User: user, Text: text, Date: &timeNow}, nil
}

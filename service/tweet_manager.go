package service

import (
	"fmt"

	"github.com/cursoGO/domain"
)

var tweet *domain.Tweet

// PublishTweet modifica la variable Tweet
func PublishTweet(t *domain.Tweet) error {
	if t.Text == "" {
		var err error = fmt.Errorf("text is required")
		return err
	}
	if t.User == nil {
		var err error = fmt.Errorf("user is required")
		return err
	}
	tweet = t
	return nil
}

// GetTweet devuelve el tweet
func GetTweet() *domain.Tweet {
	return tweet
}

package service_test

import (
	"testing"

	"github.com/cursoGO/domain"
	"github.com/cursoGO/service"
)

func TestPubishedTweetIsSaved(t *testing.T) {
	var tweet *domain.Tweet
	var user *domain.User
	var err error

	user, err = domain.NewUser("grupoesfera")
	var text string = "This is my first tweet"

	tweet, err = domain.NewTweet(user, text)

	if err != nil {
		t.Error("NO Expected" + err.Error())
	}

	err = service.PublishTweet(tweet)

	if err != nil {
		t.Error("NO Expected" + err.Error())
	}

	var publishedTweet *domain.Tweet = service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Error("Expected tweet is", tweet)
	}

	if publishedTweet.Date == nil {
		t.Error("Date is nil")
	}
}

func TestPubishedErrorCreatingTweet(t *testing.T) {
	var tweet *domain.Tweet
	var user *domain.User
	var err error

	user, err = domain.NewUser("grupoesfera")
	var text string = ""

	tweet, err = domain.NewTweet(user, text)

	if err == nil || tweet != nil {
		t.Error("Expected error, invalid text")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

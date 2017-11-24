package service_test

import (
    "testing"

    "github.com/cursoGO/service"
)

func TestPubishedTweetIsSaved(t *testing.T){

    var tweet string = "This is my first tweet"

    service.PublishTweet(tweet)

    if service.Tweet != tweet {
        t.Error("Expected tweet is", tweet)
    }

}
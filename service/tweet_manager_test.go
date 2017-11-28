package service_test

import (
	"testing"

	"github.com/cursoGO/domain"
	"github.com/cursoGO/service"
)

func TestPubishedTweetIsSaved(t *testing.T) {

	var err error

	service.RegistrarUsuario("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
	service.Login("lurome_96@hotmail.com", "123456")

	var text string = "This is my first tweet"

	err = service.PublishTweet(text, "luuliirome")

	if err != nil {
		t.Error("NO Expected" + err.Error())
		return
	}

	var publishedTweet *domain.Tweet
	publishedTweet, err = service.GetLastTweet()

	if publishedTweet.User.Nickname != "luuliirome" &&
		publishedTweet.Text != text {
		t.Error("Expected tweet is luuliirome: " + text)
		return
	}

	if publishedTweet.Date == nil {
		t.Error("Date is nil")
		return
	}
}

func TestPubishedErrorCreatingTweet(t *testing.T) {
	var tweet *domain.Tweet
	var user *domain.User
	var err error

	user, err = domain.NewUser("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
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

func TestPublishMultipleTweets(t *testing.T) {

	var user *domain.User
	var err error

	user, err = domain.NewUser("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
	verifyError(err, t)

	service.RegistrarUsuario("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
	service.Login("luuliirome", "123456")

	var text1 string = " tuvieja 1"
	var text2 string = "lucia gata"

	err = service.PublishTweet(text1, user.Nickname)
	verifyError(err, t)
	err = service.PublishTweet(text2, user.Nickname)
	verifyError(err, t)

	cantidadTweets, _ := service.CantidadDeTweets(user.Nickname)

	if cantidadTweets != 2 {
		t.Error("La cantidad de tweets publicados por el user es incorrecta")
	}

}

func TestGetByID(t *testing.T) {

	var user *domain.User
	var err error

	user, err = domain.NewUser("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")

	service.RegistrarUsuario("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
	service.Login("luuliirome", "123456")

	var text string = "Gregorio puto"

	_ = service.PublishTweet(text, user.Nickname)

	var tweet *domain.Tweet
	tweet, err = service.GetTweetByID(1)

	if err != nil {
		t.Error(err.Error())
	}

	if tweet.Text != text {
		t.Error("")
	}

}

func TestLogout(t *testing.T) {
	service.RegistrarUsuario("Lucia", "lurome_96@hotmail.com", "luuliirome", "123456")
	service.Login("luuliirome", "123456")

	var text string = "Gregorio puto"

	_ = service.PublishTweet(text, "luuliirome")

	_ = service.Logout("luuliirome", "123456")

	var err error
	err = service.PublishTweet(text, "luuliirome")

	if err == nil {
		t.Error("Expected error")
	}

}

func verifyError(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
	}
}

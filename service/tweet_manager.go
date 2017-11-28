package service

import (
	"fmt"

	"github.com/cursoGO/domain"
)

var tweets []*domain.Tweet
var usuarios []*domain.User
var logueados []*domain.User

func Login(userID string, contraseña string) error {

	for i := range usuarios {
		if (userID == usuarios[i].Nickname || userID == usuarios[i].Mail) && contraseña == usuarios[i].Contraseña {
			logueados = append(logueados, usuarios[i])
			return nil
		}
	}

	var err error = fmt.Errorf("user is required")
	return err
}

// PublishTweet modifica la variable Tweet
func PublishTweet(text string, userID string) error {

	var err error
	var user *domain.User

	if text == "" {
		err = fmt.Errorf("text is required")
		return err
	}

	for i := range logueados {
		if logueados[i].Nickname == userID || logueados[i].Mail == userID {
			user = logueados[i]
			break
		}
	}

	if user == nil {
		return fmt.Errorf("Usuario NO logueado")
	}

	var tweet *domain.Tweet

	tweet, err = domain.NewTweet(user, text)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	tweets = append(tweets, tweet)
	user.Tweets = append(user.Tweets, tweet)
	return nil

}

func RegistrarUsuario(name string, mail string, nick string, contraseña string) error {

	for i := range usuarios {
		if mail == usuarios[i].Mail || nick == usuarios[i].Nickname {
			return fmt.Errorf("Usuario ya existente")
		}
	}

	var user *domain.User
	var err error

	user, err = domain.NewUser(name, mail, nick, contraseña)

	if err != nil {
		return err
	}

	usuarios = append(usuarios, user)
	return nil
}

// GetTweet devuelve el tweet
func GetLastTweet() (*domain.Tweet, error) {
	if len(tweets) == 0 {
		var err = fmt.Errorf("No hay tweets publicados")
		return nil, err
	}
	return tweets[len(tweets)-1], nil
}

func GetTweetByID(id int) (*domain.Tweet, error) {

	for i := range tweets {
		if tweets[i].Id == id {
			return tweets[i], nil
		}
	}

	var err error = fmt.Errorf("ID invalido")

	return nil, err

}

func CantidadDeTweets(userID string) (int, error) {

	var user *domain.User
	var err error

	user, err = GetUserByID(userID)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return 0, err
	}

	return len(user.Tweets), nil
}

func GetUserByID(userID string) (*domain.User, error) {

	for i := range usuarios {
		if userID == usuarios[i].Mail || userID == usuarios[i].Nickname {
			return usuarios[i], nil
		}
	}
	var err error = fmt.Errorf("Usuario inexistente")
	return nil, err
}

func Follow(userID1 string, userID2 string) error {

	var user *domain.User
	var err error

	user, err = GetUserByID(userID1)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return err
	}

	var follower *domain.User

	follower, err = GetUserByID(userID2)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return err
	}

	user.Followers = append(user.Followers, follower)
	follower.Following = append(follower.Followers, user)

	return nil

}

func GetTweetsByUser(userID string) ([]*domain.Tweet, error) {

	var user *domain.User
	var err error

	user, err = GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user.Tweets, nil
}

func Logout(userID string, pass string) error {

	for i := range logueados {
		if (userID == logueados[i].Nickname || userID == logueados[i].Mail) && pass == logueados[i].Contraseña {
			logueados = append(logueados[:i], logueados[i+1:]...)
			return nil
		}
	}
	var err error = fmt.Errorf("Usuario no logueado")
	return err
}

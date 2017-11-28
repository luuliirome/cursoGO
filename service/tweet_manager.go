package service

import (
	"fmt"

	"github.com/cursoGO/domain"
)

type TweetManager struct {
	tweets    []*domain.Tweet
	usuarios  []*domain.User
	logueados []*domain.User
}

func (tm *TweetManager) Login(userID string, contraseña string) error {

	for i := range tm.usuarios {
		if (userID == tm.usuarios[i].Nickname || userID == tm.usuarios[i].Mail) && contraseña == tm.usuarios[i].Contraseña {
			tm.logueados = append(tm.logueados, tm.usuarios[i])
			return nil
		}
	}

	var err error = fmt.Errorf("user is required")
	return err
}

// PublishTweet modifica la variable Tweet
func (tm *TweetManager) PublishTweet(text string, userID string) error {

	var err error
	var user *domain.User

	if text == "" {
		err = fmt.Errorf("text is required")
		return err
	}

	for i := range tm.logueados {
		if tm.logueados[i].Nickname == userID || tm.logueados[i].Mail == userID {
			user = tm.logueados[i]
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

	tm.tweets = append(tm.tweets, tweet)
	user.Tweets = append(user.Tweets, tweet)
	return nil

}

func (tm *TweetManager) RegistrarUsuario(name string, mail string, nick string, contraseña string) error {

	for i := range tm.usuarios {
		if mail == tm.usuarios[i].Mail || nick == tm.usuarios[i].Nickname {
			return fmt.Errorf("Usuario ya existente")
		}
	}

	var user *domain.User
	var err error

	user, err = domain.NewUser(name, mail, nick, contraseña)

	if err != nil {
		return err
	}

	tm.usuarios = append(tm.usuarios, user)
	return nil
}

// GetTweet devuelve el tweet
func (tm *TweetManager) GetLastTweet() (*domain.Tweet, error) {
	if len(tm.tweets) == 0 {
		var err = fmt.Errorf("No hay tweets publicados")
		return nil, err
	}
	return tm.tweets[len(tm.tweets)-1], nil
}

func (tm *TweetManager) GetTweetByID(id int) (*domain.Tweet, error) {

	for i := range tm.tweets {
		if tm.tweets[i].Id == id {
			return tm.tweets[i], nil
		}
	}

	var err error = fmt.Errorf("ID invalido")

	return nil, err

}

func (tm *TweetManager) CantidadDeTweets(userID string) (int, error) {

	var user *domain.User
	var err error

	user, err = tm.GetUserByID(userID)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return 0, err
	}

	return len(user.Tweets), nil
}

func (tm *TweetManager) GetUserByID(userID string) (*domain.User, error) {

	for i := range tm.usuarios {
		if userID == tm.usuarios[i].Mail || userID == tm.usuarios[i].Nickname {
			return tm.usuarios[i], nil
		}
	}
	var err error = fmt.Errorf("Usuario inexistente")
	return nil, err
}

func (tm *TweetManager) Follow(userID1 string, userID2 string) error {

	var user *domain.User
	var err error

	user, err = tm.GetUserByID(userID1)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return err
	}

	var follower *domain.User

	follower, err = tm.GetUserByID(userID2)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return err
	}

	user.Followers = append(user.Followers, follower)
	follower.Following = append(follower.Followers, user)

	return nil

}

func (tm *TweetManager) GetTweetsByUser(userID string) ([]*domain.Tweet, error) {

	var user *domain.User
	var err error

	user, err = tm.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	return user.Tweets, nil
}

func (tm *TweetManager) Logout(userID string, pass string) error {

	for i := range tm.logueados {
		if (userID == tm.logueados[i].Nickname || userID == tm.logueados[i].Mail) && pass == tm.logueados[i].Contraseña {
			tm.logueados = append(tm.logueados[:i], tm.logueados[i+1:]...)
			return nil
		}
	}
	var err error = fmt.Errorf("Usuario no logueado")
	return err
}

func (tm *TweetManager) DeleteTweet(tweetId int) {

}

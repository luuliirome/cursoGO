package service

import (
	"fmt"

	"github.com/cursoGO/domain"
)

type TweetManager struct {
	tweets    map[int]domain.Tweet
	usuarios  []domain.User
	logueados map[string]*domain.User
}

func (tm *TweetManager) Login(userID string, contraseña string) error {

	for i := range tm.usuarios {
		if userID == tm.usuarios[i].Mail || userID == tm.usuarios[i].Nickname {
			if tm.usuarios[i].Contraseña == contraseña {
				tm.logueados[userID] = &tm.usuarios[i]
				return nil
			}

			return fmt.Errorf("Contraseña invalida")
		}
	}

	return fmt.Errorf("El usuario no esta registrado")

}

// PublishTweet modifica la variable Tweet
func (tm *TweetManager) PublishTweet(text string, userID string) error {

	var err error

	if text == "" {
		err = fmt.Errorf("text is required")
		return err
	}

	i, ok := tm.logueados[userID]

	if !ok {
		return fmt.Errorf("Usuario NO logueado")
	}

	var tweet domain.Tweet

	tweet, err = domain.NewTweet(i, text)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	tm.tweets[tweet.Id] = tweet
	return nil

}

func (tm *TweetManager) RegistrarUsuario(name string, mail string, nick string, contraseña string) error {

	for i := range tm.usuarios {
		if mail == tm.usuarios[i].Mail || nick == tm.usuarios[i].Nickname {
			return fmt.Errorf("Usuario ya existente")
		}
	}

	var user domain.User
	var err error

	user, err = domain.NewUser(name, mail, nick, contraseña)

	if err != nil {
		return err
	}

	tm.usuarios = append(tm.usuarios, user)
	return nil
}

func (tm *TweetManager) GetTweetByID(id int) (*domain.Tweet, error) {

	i, ok := tm.tweets[id]

	if !ok {
		var err error = fmt.Errorf("ID invalido")
		return nil, err
	}

	return &i, nil

}

// func (tm *TweetManager) CantidadDeTweets(userID string) (int, error) {

// 	var user *domain.User
// 	var err error

// 	user, err = tm.GetUserByID(userID)

// 	if err != nil {
// 		err = fmt.Errorf("Usuario no registrado")
// 		return 0, err
// 	}

// 	return len(user.Tweets), nil
// }

func (tm *TweetManager) GetUserByID(userID string) (*domain.User, error) {

	for i := range tm.usuarios {
		if userID == tm.usuarios[i].Mail || userID == tm.usuarios[i].Nickname {
			return &tm.usuarios[i], nil
		}
	}
	var err error = fmt.Errorf("Usuario inexistente")
	return nil, err
}

func (tm *TweetManager) Follow(seguido string, seguidor string) error {

	var user *domain.User
	var err error

	user, err = tm.GetUserByID(seguido)

	if err != nil {
		err = fmt.Errorf("Usuario no registrado")
		return err
	}

	var follower *domain.User

	for i := range tm.logueados {
		if tm.logueados[i].Nickname == seguidor || tm.logueados[i].Mail == seguidor {
			user = tm.logueados[i]
			break
		}
	}

	if user == nil {
		return fmt.Errorf("Usuario NO logueado")
	}

	user.Followers = append(user.Followers, follower)
	follower.Following = append(follower.Followers, user)

	return nil

}

func (tm *TweetManager) GetTweetsByUser(userID string) ([]domain.Tweet, error) {

	var user *domain.User
	var err error

	user, err = tm.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	var userTweets []domain.Tweet

	for _, v := range tm.tweets {
		if v.User == user {
			userTweets = append(userTweets, v)
		}
	}

	if len(userTweets) == 0 {
		err = fmt.Errorf("El usuario no tiene tweets publicados")
		return nil, err
	}

	return userTweets, nil
}

func (tm *TweetManager) Logout(userID string, pass string) error {

	i, ok := tm.logueados[userID]

	if !ok {
		return fmt.Errorf("Usuario no logueado")
	}

	if i.Contraseña != pass {
		return fmt.Errorf("Contraseña incorrecta")
	}

	delete(tm.logueados, userID)
	return nil
}

func (tm *TweetManager) DeleteTweet(tweetId int, userID string, pass string) error {

	i, ok := tm.logueados[userID]

	if !ok {
		return fmt.Errorf("Usuario no logueado")
	}

	if i.Contraseña != pass {
		return fmt.Errorf("Contraseña incorrecta")
	}

	v, ok := tm.tweets[tweetId]

	if !ok {
		return fmt.Errorf("Tweet inexistente")
	}

	if v.User.Nickname != i.Nickname {
		return fmt.Errorf("No posees los permisos necesarios para borrar este tweet")
	}

	delete(tm.tweets, tweetId)
	return nil
}

func (tm *TweetManager) EditTweet(tweetId int, newTweet string, userID string) error {

	_, ok := tm.logueados[userID]

	if !ok {
		return fmt.Errorf("Usuario no logueado")
	}

	t, ok := tm.tweets[tweetId]

	if !ok {
		return fmt.Errorf("Tweeter inexistente")
	}

	if t.User.Nickname != userID {
		return fmt.Errorf("No posees los permisos necesarios para modificar el tweet")
	}

	tweet, err := domain.NewTweet(t.User, newTweet)

	if err != nil {
		return err
	}

	tm.tweets[tweetId] = tweet

	return nil
}

func (tm *TweetManager) ShowFollowers(userID string) ([]*domain.User, error) {
	var user *domain.User
	var err error

	user, err = tm.GetUserByID(userID)

	if err != nil {
		return nil, err
	}

	if len(user.Followers) == 0 {
		return nil, fmt.Errorf("El usuario no posee seguidores")
	}

	return user.Followers, nil
}

func (tm *TweetManager) ShowTimeline(userID string) ([]*domain.Tweet, error) {

	var logueado bool = false
	for j := range tm.logueados {
		if tm.logueados[j].Nickname == userID || tm.logueados[j].Mail == userID {
			logueado = true
			break
		}
	}

	if !logueado {
		return nil, fmt.Errorf("Usuario no logueado")
	}

	return nil, nil
}

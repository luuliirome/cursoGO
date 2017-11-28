package main

import (
	"github.com/abiosoft/ishell"
	"github.com/cursoGO/domain"
	"github.com/cursoGO/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			if text == "" {
				c.Print("Texto vacio!\n")
				return
			}

			c.Print("Write your nick or email: ")

			userID := c.ReadLine()

			var err error = service.PublishTweet(text, userID)

			if err != nil {
				c.Print(err.Error() + "\n")
				return
			}

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			var err error

			tweet, err := service.GetLastTweet()

			if err != nil {
				c.Print(err.Error())
			}

			c.Println(tweet.User.Name + ": " + tweet.Text + " - Fecha: " + tweet.Date.String())

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "Te loguea",
		Func: func(c *ishell.Context) {

			var err error

			defer c.ShowPrompt(true)

			c.Print("Enter your nick or email: ")

			name := c.ReadLine()

			c.Print("Enter your password: ")

			pass := c.ReadLine()

			err = service.Login(name, pass)

			if err != nil {
				c.Print(err.Error())
				return
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Enter your name: ")

			name := c.ReadLine()

			c.Print("Enter your mail: ")

			mail := c.ReadLine()

			c.Print("Enter your nickname: ")

			nick := c.ReadLine()

			c.Print("Enter your password: ")

			pass := c.ReadLine()

			var err error = service.RegistrarUsuario(name, mail, nick, pass)

			if err != nil {
				c.Print(err.Error())
				c.Print("\n")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAmountOfTweets",
		Help: "Shows the amount of tweets publish by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Enter your nickname or mail: ")

			userID := c.ReadLine()
			var cant int
			var err error

			cant, err = service.CantidadDeTweets(userID)

			if err != nil {
				c.Print(err.Error())
			}
			c.Print(cant)
			c.Print("\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Logout",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Enter your nickname or mail: ")
			userID := c.ReadLine()

			c.Print("Enter your password: ")
			pass := c.ReadLine()

			_ = service.Logout(userID, pass)

		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "GetAllMyTweets",
		Help: "Shows tweets publish by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Enter your nick or mail: ")

			userID := c.ReadLine()
			var tweets []*domain.Tweet
			var err error

			tweets, err = service.GetTweetsByUser(userID)

			if err != nil {
				c.Print(err.Error())
			}

			for i := range tweets {
				c.Print(tweets[i].User.Name + ": " + tweets[i].Text + " - Fecha: " + tweets[i].Date.String() + "\n")
			}
			c.Print("\n")

			return
		},
	})

	shell.Run()

}

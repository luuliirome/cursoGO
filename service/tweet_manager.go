package service

// Tweet es la variable que guarda el tweet actual
var Tweet string 

// PublishTweet modifica a variable Tweet
func PublishTweet(tweet string){
	Tweet = tweet
}
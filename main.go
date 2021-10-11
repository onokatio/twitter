package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
)

func main(){
	godotenv.Load()
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMERKEY"), os.Getenv("TWITTER_CONSUMERSECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESSTOKEN"), os.Getenv("TWITTER_ACCESSSECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	    Count: 1,
	})
	if err != nil {
		pp.Println(err)
	}
	pp.Println(tweets)
}

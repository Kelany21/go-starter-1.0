package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

var stream *twitter.Stream = nil

func main() {
	/**
	* start multi language
	 */
	//err := gotrans.InitLocales("public/trans")
	//if err != nil {
	//	panic(err)
	//}
	///**
	//* add custom role to validation
	//*
	//*/
	//requests.Init()
	///**
	//* connect with data base logic you can edit .env file to
	//* change any connection params
	// */
	//config.ConnectToDatabase()
	///**
	//* drop All tables and migrate
	//* to stop delete tables make DROP_ALL_TABLES false in env file
	//* if you need to stop auto migration just stop this line
	//*/
	//models.MigrateAllTable(os.Getenv("PRODUCTION_MODEL_PATH"))
	///**
	//* this function will open seeders folder look inside all files
	//* search for seeders function and seed execute these function
	//* if you need to stop seeding you can stop this line
	//*/
	//seeders.Seed()

	creds := Credentials{
		AccessToken:       "1094561861540999169-q6T9hun6c9ToqKE273cdv2VhfqgrEM",
		AccessTokenSecret: "yTohYgoLm7MH4T1Pznuzn4ptVajNAaOHrIwlQZKH8VKk9",
		ConsumerKey:       "EwSQRgOThbJZWLq728Ms5elap",
		ConsumerSecret:    "g61LASw56nYKXgHU22hibySDoSaHLsNM5mwBaA9Gq1uX6tHIqG",
	}
	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	//params := &twitter.StreamFilterParams{
	//	Track:         []string{"omar_kelany98"},
	//	StallWarnings: twitter.Bool(true),
	//}
	//stream, _ = client.Streams.Filter(params)
	//demux := twitter.NewSwitchDemux()
	//demux.Tweet = func(tweet *twitter.Tweet) {
	//	fmt.Println(123)
	//	//go createTwitter(tweet)
	//}
	//demux.HandleChan(stream.Messages)
	var ids []int64
	ids = append(ids, 1258106943153098754)
	search, _, _ := client.Statuses.Lookup(ids, &twitter.StatusLookupParams{})
	//users, _, _ := client.Users.Lookup(&twitter.UserLookupParams{
	//	UserID:          nil,
	//	ScreenName:      []string{"omar_kelany98", "OmarAme20770485"},
	//	IncludeEntities: nil,
	//})
	//search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
	//	Query:           "(from:omar_kelany98)",
	//	Geocode:         "",
	//	Lang:            "",
	//	Locale:          "",
	//	ResultType:      "",
	//	Count:           100,
	//	SinceID:         0,
	//	MaxID:           0,
	//	Until:           "",
	//	Since:           "",
	//	Filter:          "",
	//	IncludeEntities: nil,
	//	TweetMode:       "",
	//})

	//result := make(map[string]string)
	//op := false
	//count := 0
	//searchT, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
	//	Query: "(from:ZHindustanTamil)",
	//	Count: 100,
	//})
	//for i := 0; i < len(searchT.Statuses); i++ {
	//	_, ok := result[searchT.Statuses[i].IDStr]
	//	if !ok {
	//		result[searchT.Statuses[i].IDStr] = searchT.Statuses[i].IDStr
	//	}
	//}
	//fmt.Println("search: ", len(searchT.Statuses))
	//fmt.Println("end search: ", len(result))
	//var id int64
	//if len(searchT.Statuses) != 0 {
	//	op = true
	//	id = searchT.Statuses[len(searchT.Statuses)-1].ID
	//}
	//for op {
	//	if count == 1000 {
	//		fmt.Println("end search: ", len(result))
	//		return
	//	}
	//	searchT, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
	//		Query: "(from:ZHindustanTamil)",
	//		Count: 1000,
	//		MaxID: id,
	//	})
	//	if len(searchT.Statuses) <= 1 {
	//		op = false
	//	} else {
	//		count += len(searchT.Statuses)
	//		for i := 0; i < len(searchT.Statuses); i++ {
	//				_, ok := result[searchT.Statuses[i].IDStr]
	//				if !ok {
	//					result[searchT.Statuses[i].IDStr] = searchT.Statuses[i].IDStr
	//				}
	//		}
	//		fmt.Println("search: ", len(searchT.Statuses))
	//		fmt.Println("end search: ", len(result))
	//		id = searchT.Statuses[len(searchT.Statuses)-1].ID
	//	}
	//}
	//
	////if err != nil {
	////	log.Print(err)
	////}
	////fmt.Println("resp: ",resp)
	//fmt.Println("end2 search: ", len(result))
	//fmt.Println("end2 search: ", id)
	//fmt.Println("tweet: ", strings.ReplaceAll(searchT.Statuses[len(searchT.Statuses)-1].User.ProfileImageURLHttps, "_normal", ""))
	for _, user := range search[0].Entities.UserMentions{
		fmt.Println("hashtag text: ", user.ScreenName)
	}
	//fmt.Println("search: ", search.Statuses)
	//fmt.Println("search: ",search[1].User)
	//fmt.Println("search 1 Likes: ",search[0].FavoriteCount)
	//fmt.Println("search 1 Retweets: ",search[0].RetweetCount)
	//fmt.Println("search 1 Replies: ",search[0].ReplyCount)
	//fmt.Println("search 2 Likes: ",search[1].FavoriteCount)
	//fmt.Println("search 2 Retweets: ",search[1].RetweetCount)
	//fmt.Println("search 3 Replies: ",search[1].ReplyCount)
	//fmt.Println("---------------------------------")
	//fmt.Println("user: ",users[0])
	//fmt.Println("user: ",users[1])
	//fmt.Println("search Likes: ",search[1].Place.Country)
	//fmt.Println("search first: ",search.Statuses[len(search.Statuses)-1])
	//fmt.Println("search len: ",len(search.Statuses))
	/**
	* Run gin framework
	* add middleware
	* run routing
	* serve app
	 */
	//providers.Run()
}

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

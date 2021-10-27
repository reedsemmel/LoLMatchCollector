package main

import (
	"context"
	"log"
	"os"

	lol "github.com/reedsemmel/LoLMatchCollector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client lol.ApiClient
	ec     chan error
	coll   *mongo.Collection
)

func insert_match(match *lol.MatchDto) error {
	buf, err := bson.Marshal(*match)
	if err != nil {
		return err
	}
	_, err = coll.InsertOne(context.Background(), buf)
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return err
	}
	return nil
}

func collect_page(r lol.Rank, d lol.Division, page int) {
	people, err := client.GetLeagueEntries(r, d, page)
	if err != nil {
		ec <- err
		return
	}
	for _, person := range people {
		puuid, err := client.SummonerIdToPuuid(person.SummonerId)
		if err != nil {
			ec <- err
			continue
		}

		matches, err := client.GetMatchesForPuuid(puuid)
		if err != nil {
			ec <- err
			continue
		}
		for _, matchId := range matches {
			// TODO: Check to see if the match is in the DB before making the API call
			match, err := client.GetMatch(matchId)
			if err != nil {
				ec <- err
				continue
			}
			match.MatchId = matchId
		}
	}
}

func main() {
	apiKey := os.Getenv("RGAPIKEY")
	mongoUrl := os.Getenv("MONGODB_ADDR")
	user := os.Getenv("MONGODB_USER")
	password := os.Getenv("MONGODB_PASS")

	mongoOpts := options.Client().ApplyURI(mongoUrl)
	mongoOpts.SetAppName("LolMatchCollector")
	mongoOpts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    "admin",
		Username:      user,
		Password:      password,
	})

	if apiKey == "" {
		log.Fatalln("No $RGAPIKEY found")
	}

	mongoClient, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	coll = mongoClient.Database("admin").Collection("matches")
	if coll == nil {
		log.Fatal("could not find collection")
	}

	client = lol.NewApiClient(apiKey)
	ec = make(chan error)

}

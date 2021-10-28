package main

import (
	"context"
	"encoding/json"
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

func insertMatch(matchId string) error {
	res := coll.FindOne(context.Background(), bson.M{"matchid": matchId}, options.FindOne().SetProjection(bson.M{"matchid": 1, "_id": 0}))

	if res.Err() == mongo.ErrNoDocuments {
		match, err := client.GetMatch(matchId)
		if err != nil {
			ec <- err
			return nil
		}
		buf, err := bson.Marshal(match)
		if err != nil {
			return err
		}
		_, err = coll.InsertOne(context.Background(), buf)
		if err != nil && !mongo.IsDuplicateKeyError(err) {
			return err
		}
		return nil
	}

	// This should be nil if the document was found. In that case we don't want to do anything
	// and we are good.
	return res.Err()
}

func collectPage(r lol.Rank, d lol.Division, page int) {
	people, err := client.GetLeagueEntries(r, d, page)
	if err != nil {
		log.Fatal(err)
	}
	for i, person := range people {
		log.Printf("[%v/%v]: %v\n", i, len(people), person.SummonerId)
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
			err = insertMatch(matchId)
			if err != nil {
				// Database errors should not happen. This should be fatal
				log.Fatal(err)
			}
		}
	}
}

type Task struct {
	Rank int
	Div  int
	Page int
}

func do_task() {
	start := -1

	dat, err := os.ReadFile("tasks.json")

	if err != nil {
		log.Fatalln("cannot ReadFile tasks.json:", err)
	}

	var tasks []Task
	var current_task Task

	err = json.Unmarshal(dat, &tasks)
	if err != nil {
		log.Fatalln("cannot unmarshal tasks.json:", err)
	}

	dat, err = os.ReadFile("current_task.json")
	if err != nil {
		// If it doesn't exist we start at the beginning
		start = 0
	} else {
		err = json.Unmarshal(dat, &current_task)
		if err != nil {
			log.Fatalln("cannot unmarshal current_task.json:", err)
		}
		// Find where to start
		for i, v := range tasks {
			if v.Div == current_task.Div && v.Rank == current_task.Rank && v.Page == current_task.Page {
				start = i
				break
			}
		}
	}
	if start == -1 {
		log.Fatalln("could not find where to start")
	}

	for i := start; i < len(tasks); i++ {
		log.Printf("*** TASK [%v/%v] ***", i, len(tasks))
		// Write current task to current_task.json file
		dat, err = json.Marshal(tasks[i])
		if err != nil {
			log.Fatal("failed to marshal the current task")
		}
		err = os.WriteFile("current_task.json", dat, 0644)
		if err != nil {
			log.Fatal("WriteFile:", err)
		}

		collectPage(lol.Rank(tasks[i].Rank), lol.Division(tasks[i].Div), tasks[i].Page)
	}

}

func main() {
	log.Println("===== Collector started =====")

	apiKey := os.Getenv("RGAPIKEY")
	mongoUrl := os.Getenv("MONGODB_ADDR")
	user := os.Getenv("MONGODB_USER")
	password := os.Getenv("MONGODB_PASS")
	collection_name := os.Getenv("MONGODB_COLLECTION")

	if len(apiKey) == 0 || len(mongoUrl) == 0 || len(user) == 0 || len(password) == 0 || len(collection_name) == 0 {
		log.Fatal("not all needed environment variables are set")
	}

	mongoOpts := options.Client().ApplyURI(mongoUrl)
	mongoOpts.SetAppName("LolMatchCollector")
	mongoOpts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    "admin",
		Username:      user,
		Password:      password,
	})

	mongoClient, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	coll = mongoClient.Database("admin").Collection(collection_name)

	if coll == nil {
		log.Fatal("could not find collection")
	}

	client = lol.NewApiClient(apiKey)
	ec = make(chan error)

	go error_handler()

	do_task()
}

func error_handler() {
	i := 0
	for {
		err := <-ec
		i++
		log.Printf("received error \"%v\", %v/5 before stopping\n", err, i)
		if i == 5 {
			log.Fatalln("too many errors. stopping.")
		}
	}
}

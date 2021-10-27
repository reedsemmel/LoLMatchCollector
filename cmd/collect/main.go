package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	lol "github.com/reedsemmel/LoLMatchCollector"
)

func main() {
	apiKey := os.Getenv("RGAPIKEY")

	if apiKey == "" {
		log.Fatalln("No $RGAPIKEY found")
	}

	client := lol.NewApiClient(apiKey)

	players, err := client.GetLeagueEntries(lol.RankG, lol.Div4, 1)
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range players {
		puuid, err := client.SummonerIdToPuuid(p.SummonerId)
		if err != nil {
			log.Fatalln(err)
		}
		matches, err := client.GetMatchesForPuuid(puuid)
		if err != nil {
			log.Fatalln(err)
		}
		for _, match := range matches {
			res, err := client.GetMatch(match)
			if err != nil {
				log.Fatalln(err)
			}
			buf, err := json.Marshal(res)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(buf))
			os.Exit(0)
		}
	}
}

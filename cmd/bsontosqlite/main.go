package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MatchDto struct {
	Metadata MetadataDto `json:"metadata"`
	Info     InfoDto     `json:"info"`
	MatchId  string      `json:"matchId"`
}

type MetadataDto struct {
	DataVersion string `json:"dataVersion"`
	MatchId     string `json:"matchId"`
}

type InfoDto struct {
	GameVersion  string           `json:"gameVersion"`
	Participants []ParticipantDto `json:"participants"`
}

type ParticipantDto struct {
	Assists                        int    `json:"assists"`
	ChampionId                     int    `json:"championId"`
	ChampionName                   string `json:"championName"`
	ChampionTransform              int    `json:"championTransform"`
	Deaths                         int    `json:"deaths"`
	GoldEarned                     int    `json:"goldEarned"`
	GoldSpent                      int    `json:"goldSpent"`
	Item0                          int    `json:"item0"`
	Item1                          int    `json:"item1"`
	Item2                          int    `json:"item2"`
	Item3                          int    `json:"item3"`
	Item4                          int    `json:"item4"`
	Item5                          int    `json:"item5"`
	Item6                          int    `json:"item6"`
	ItemsPurchased                 int    `json:"itemsPurchased"`
	Kills                          int    `json:"kills"`
	Lane                           string `json:"lane"`
	NeutralMinionsKilled           int    `json:"neutralMinionsKilled"`
	ParticipantId                  int    `json:"participantId"`
	PentaKills                     int    `json:"pentaKills"`
	Spell1Casts                    int    `json:"spell1Casts"`
	Spell2Casts                    int    `json:"spell2Casts"`
	Spell3Casts                    int    `json:"spell3Casts"`
	Spell4Casts                    int    `json:"spell4Casts"`
	SummonerName                   string `json:"summonerName"`
	TeamId                         int    `json:"teamId"`
	TeamPosition                   string `json:"teamPosition"`
	TotalDamageDealt               int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int    `json:"totalDamageTaken"`
	TotalHeal                      int    `json:"totalHeal"`
	TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int    `json:"totalUnitsHealed"`
	TrueDamageDealt                int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int    `json:"trueDamageTaken"`
	UnrealKills                    int    `json:"unrealKills"`
	VisionScore                    int    `json:"visionScore"`
	Win                            bool   `json:"win"`

	MatchId string `json:"matchId"`
	Patch   string `json:"patch"`
}

func main() {
	/*itemData, err := os.ReadFile("item.json")
	if err != nil {
		log.Fatal(err)
	}
	var itemMap map[string]interface{}
	if err := json.Unmarshal(itemData, &itemMap); err != nil {
		log.Fatal(err)
	}
	itemMapData := itemMap["data"].(map[string]interface{})
	for k, v := range itemMapData {
		name := v.(map[string]interface{})["name"].(string)
		fmt.Println(k, name)
	}*/

	mongoOpts := options.Client().ApplyURI("mongodb://localhost:27017")

	mongoClient, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	coll := mongoClient.Database("admin").Collection("matches")

	cur, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	count := 0
	for cur.Next(context.Background()) {
		count++
		var result MatchDto
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		if count%1000 == 0 {
			log.Println(count)
		}
		for _, p := range result.Info.Participants {
			p.MatchId = result.MatchId
			p.Patch = result.Info.GameVersion
			js, err := json.Marshal(p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", js)
		}
	}
}

package lol

const (
	queueType = "RANKED_SOLO_5x5"
)

type Rank int
type Division int

const (
	RankChall Rank = iota
	RankGM
	RankM
	RankD
	RankP
	RankG
	RankS
	RankB
	RankI

	Div1 Division = iota
	Div2
	Div3
	Div4
)

var Ranks = map[Rank]string{
	RankChall: "CHALLENGER",
	RankGM:    "GRANDMASTER",
	RankM:     "MASTER",
	RankD:     "DIAMOND",
	RankP:     "PLATINUM",
	RankG:     "GOLD",
	RankS:     "SILVER",
	RankB:     "BRONZE",
	RankI:     "IRON",
}

var Divisions = map[Division]string{
	Div1: "I",
	Div2: "II",
	Div3: "III",
	Div4: "IV",
}

type LeagueEntryDto struct {
	SummonerId   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type MatchDto struct {
	Metadata MetadataDto `json:"metadata"`
	Info     InfoDto     `json:"info"`
	MatchId  string      `json:"_id"` // Set the default key to be the match id
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
}

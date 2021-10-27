package lol

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/time/rate"
)

type ApiClient struct {
	apiKey   string
	limiter  *rate.Limiter
	platform string
	region   string
}

// Creates a new api client using the provided api key
func NewApiClient(apiKey string) ApiClient {
	// Limits us to 96 requests every 2 minutes, just 4 below the bucket allows
	limiter := rate.NewLimiter(0.8, 1)
	return ApiClient{
		apiKey:   apiKey,
		limiter:  limiter,
		platform: "https://na1.api.riotgames.com",
		region:   "https://americas.api.riotgames.com",
	}
}

// rawRequest makes the actual http call, and will wait before doing so according to the limiter
func (c *ApiClient) rawRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Riot-Token", c.apiKey)

	err = c.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

// request will call rawRequest and retry on 429s to get around service-level rate limiting
func (c *ApiClient) request(server, endpoint string) (*http.Response, error) {
	for i := 1; i < 4; i++ {
		res, err := c.rawRequest(server + endpoint)
		if err != nil {
			return nil, err
		}
		switch res.StatusCode {
		case 200, 404:
			return res, err
		case 400, 401, 403, 500, 502, 503, 504:
			log.Fatalln(res.Status, server+endpoint)
		case 429:
			res.Body.Close()

			retry_after, err := strconv.Atoi(res.Header.Get("Retry-After"))
			if err != nil {
				retry_after = 15
			}

			if i == 3 {
				break
			}

			retry_after = retry_after*i + 5

			log.Println("Received a 429, waiting", retry_after, "seconds")
			time.Sleep(time.Duration(retry_after) * time.Second)
		default:
			log.Fatalln("Received unexpected status code", res.Status)
		}
	}
	log.Fatalln("too many 429s on", server+endpoint)
	return nil, nil
}

func (c *ApiClient) GetLeagueEntries(r Rank, d Division, page int) ([]LeagueEntryDto, error) {
	// Make the endpoint string
	if page < 1 {
		return nil, errors.New("page number cannot be than 1")
	}

	if (r == RankChall || r == RankGM || r == RankM) && d != Div1 {
		return nil, errors.New("apex tiers only use division 1")
	}
	endpoint := fmt.Sprintf("/lol/league-exp/v4/entries/%v/%v/%v?page=%v", queueType, Ranks[r],
		Divisions[d], page)

	res, err := c.request(c.platform, endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var players []LeagueEntryDto

	if res.StatusCode == 404 {
		return nil, nil
	}

	if err := json.Unmarshal(buf, &players); err != nil {
		return nil, err
	}

	return players, nil
}

func (c *ApiClient) SummonerIdToPuuid(summoner_id string) (string, error) {
	endpoint := fmt.Sprintf("/lol/summoner/v4/summoners/%v", summoner_id)

	res, err := c.request(c.platform, endpoint)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("invalid response %v", res.Status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var j map[string]interface{}

	if err := json.Unmarshal(buf, &j); err != nil {
		return "", err
	}
	if puuid, ok := j["puuid"]; ok {
		return fmt.Sprintf("%v", puuid), nil
	}
	return "", fmt.Errorf("json response did not contain a puuid field")
}

func (c *ApiClient) GetMatchesForPuuid(puuid string) ([]string, error) {
	// startTime is June 16th, 2021, which is when Riot started including time stamps in their
	// matches. This is convenient since we are only looking at mythics.
	// queue=420 is ranked solo/duo (nice)
	endpoint := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%v/ids?startTime=1623801600&queue=420&count=25", puuid)

	res, err := c.request(c.region, endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response %v", res.Status)
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var matches []string
	if err := json.Unmarshal(buf, &matches); err != nil {
		return nil, err
	}

	return matches, nil
}

func (c *ApiClient) GetMatch(match_id string) (MatchDto, error) {
	endpoint := fmt.Sprintf("/lol/match/v5/matches/%v", match_id)

	res, err := c.request(c.region, endpoint)
	if err != nil {
		return MatchDto{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return MatchDto{}, fmt.Errorf("bad status %v", res.Status)
	}

	buf, err := io.ReadAll(res.Body)

	if err != nil {
		return MatchDto{}, err
	}

	var match MatchDto

	if err := json.Unmarshal(buf, &match); err != nil {
		return MatchDto{}, err
	}

	return match, nil

}

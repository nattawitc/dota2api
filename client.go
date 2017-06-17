package dota2api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	matchUrl = "https://api.steampowered.com/IDOTA2Match_570"
	ecoUrl   = "https://api.steampowered.com/IEconDOTA2_570"

	getLeagueListingUrl             = "/GetLeagueListing/v001"
	getLiveLeagueGamesUrl           = "/GetLiveLeagueGames/v001"
	getMatchDetailsUrl              = "/GetMatchDetails/v001"
	getMatchHistoryUrl              = "/GetMatchHistory/v001"
	getMatchHistoryBySequenceNumUrl = "/GetMatchHistoryBySequenceNum/v001"
	getScheduledLeagueGamesUrl      = "/GetScheduledLeagueGames/v001"
	getTeamInfoByTeamIDUrl          = "/GetTeamInfoByTeamID/v001"
	getTournamentPlayerStatsUrl     = "/GetTournamentPlayerStats/v001"
	getTopLiveGameUrl               = "/GetTopLiveGame/v001"

	getGameItemsUrl            = "/GetGameItems/v001"
	getItemIconPathUrl         = "/GetItemIconPath/v001"
	getHeroesUrl               = "/GetHeroes/v001"
	getRaritiesUrl             = "/GetRarities/v001"
	getTournamentPrizePoolUrl  = "/GetTournamentPrizePool/v001"
	getEventStatsForAccountUrl = "/GetEventStatsForAccount/v001"
)

var ()

type Client struct {
	key string
}

func New(key string) Client {
	return Client{key}
}

func (c Client) GetLeagueListing() (GetLeagueListingResult, error) {
	errHeader := "GetLeagueListing:"
	url := matchUrl + getLeagueListingUrl + "?key=" + c.key
	resp, err := http.Get(url)
	if err != nil {
		return GetLeagueListingResult{}, fmt.Errorf("%v %v", errHeader, err)
	}
	switch resp.StatusCode {
	case http.StatusOK:
		res := &GetLeagueListingWrapper{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return GetLeagueListingResult{}, fmt.Errorf("%v %v", errHeader, err)
		}
		err = json.Unmarshal(body, res)
		if err != nil {
			return GetLeagueListingResult{}, fmt.Errorf("%v %v", errHeader, err)
		}
		return res.Result, nil
	case http.StatusForbidden:
		return GetLeagueListingResult{}, fmt.Errorf("%v %v", errHeader, "Access denied")
	default:
		return GetLeagueListingResult{}, fmt.Errorf("%v %v", errHeader, "Unknown error with http status", resp.StatusCode)
	}
}

func (c Client) GetLiveLeagueGames() {

}

func (c Client) GetMatchDetails() {

}

func (c Client) GetMatchHistory() {

}

func (c Client) GetMatchHistoryBySequenceNum() {

}

func (c Client) GetScheduledLeagueGames() {

}

func (c Client) GetTeamInfoByTeamID() {

}

func (c Client) GetTournamentPlayerStats() {

}

func (c Client) GetTopLiveGame() {

}

func (c Client) GetGameItems() {

}

func (c Client) GetItemIconPath() {

}

func (c Client) GetHeroes() {

}

func (c Client) GetRarities() {

}

func (c Client) GetTournamentPrizePool() {

}

func (c Client) GetEventStatsForAccount() {

}

package dota2api

type GetLeagueListingWrapper struct {
	Result GetLeagueListingResult `json:"result"`
}

type GetLeagueListingResult struct {
	League []League `json:"leagues"`
}

type League struct {
	Name          string `json:"name"`
	LeagueID      int    `json:"leagueid"`
	Description   string `json:"description"`
	TournamentURL string `json:"tournament_url"`
	ItemDef       int    `json:"itemdef"`
}

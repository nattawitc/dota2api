package dota2api

type GetLeagueListingWrapper struct {
	Result GetLeagueListingResult `json:"result"`
}

type GetLeagueListingResult struct {
	Leagues []League `json:"leagues"`
}

type League struct {
	Name          string `json:"name"`
	LeagueID      int    `json:"leagueid"`
	Description   string `json:"description"`
	TournamentURL string `json:"tournament_url"`
	ItemDef       int    `json:"itemdef"`
}

type GetLiveLeagueGamesWrapper struct {
	Result GetLiveLeagueGamesResult `json:"result"`
}

type GetLiveLeagueGamesResult struct {
	Games []Game `json:"games"`
}

type Game struct {
	Players           []Player   `json:"players"`
	RadientTeam       Team       `json:"radiant_team"`
	DireTeam          Team       `json:"dire_time"`
	LobbyID           uint64     `json:"lobby_id"`
	MatchID           uint64     `json:"match_id"`
	Spectators        int        `json:"spectators"`
	GameNumber        int        `json:"game_number"`
	LeagueID          int        `json:"league_id"`
	StreamDelayS      int        `json:"stream_delay_s"`
	RadiantSeriesWins int        `json:"radiant_series_wins"`
	DireSeriessWins   int        `json:"dire_series_wins"`
	SeriesType        int        `json:"series_type"`
	LeagueSeriesID    int        `json:"league_series_id"`
	LeagueGameID      int        `json:"league_game_id"`
	StageName         string     `json:"stage_name"`
	LeagueTier        int        `json:"league_tier"`
	Scoreboard        Scoreboard `json:"scoreboard"`
}

type Player struct {
	AccountID       int    `json:"account_id"`
	Name            string `json:"name"`
	HeroID          int    `json:"hero_id"`
	Team            int    `json:"team"`
	TeamDescription string `json:"-"`
}

type Team struct {
	TeamName string `json:"team_name"`
	TeamID   int    `json:"team_id"`
	TeamLogo int    `json:"team_logo"`
	Complete bool   `json:"complete"`
}

type Scoreboard struct {
	Duration           float64        `json:"duration"`
	RoshanRespawnTimer int            `json:"roshan_respawn_timer"`
	Radiant            ScoreboardTeam `json:"radiant"`
	Dire               ScoreboardTeam `json:"dire"`
}

type ScoreboardTeam struct {
	Score         int `json:"score"`
	TowerState    int `json:"tower_state"`
	BarracksState int `json:"barracks_state"`
	Picks         []struct {
		HeroID int `json:"hero_id"`
	} `json:"picks"`
	Bans []struct {
		HeroID int `json:"hero_id"`
	} `json:"bans"`
	Players   []ScoreboardTeamPlayer  `json:"players"`
	Abilities []ScoreboardTeamAbility `json:"abilities"`
}

type ScoreboardTeamPlayer struct {
	PlayerSlot       int     `json:"player_slot"`
	AccountID        uint64  `json:"account_id"`
	HeroID           int     `json:"hero_id"`
	Kills            int     `json:"kills"`
	Death            int     `json:"death"`
	Assists          int     `json:"assists"`
	LastHits         int     `json:"last_hits"`
	Denies           int     `json:"denies"`
	Gold             int     `json:"gold"`
	Level            int     `json:"level"`
	GoldPerMin       int     `json:"gold_per_min"`
	XPPerMin         int     `json:"xp_per_min"`
	UltimateState    int     `json:"ultimate_state"`
	UltimateCooldown int     `json:"ultimate_cooldown"`
	Item0            int     `json:"item0"`
	Item1            int     `json:"item1"`
	Item2            int     `json:"item2"`
	Item3            int     `json:"item3"`
	Item4            int     `json:"item4"`
	Item5            int     `json:"item5"`
	RespawnTimer     int     `json:"respawn_timer"`
	PositionX        float64 `json:"position_x"`
	PositionY        float64 `json:"position_y"`
	NetWorth         int     `json:"net_worth"`
}

type ScoreboardTeamAbility struct {
	AbilityID    int `json:"ability_id"`
	AbilityLevel int `json:"ability_level"`
}

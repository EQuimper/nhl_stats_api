package nhl_stats_api

type Conference struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	ShortName    string `json:"shortName"`
	Active       bool   `json:"active"`
}

type Position struct {
	Abbrev   string `json:"abbrev"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
	Type     string `json:"type"`
}

type PlayerInfo struct {
	ID                 int     `json:"id"`
	FullName           string  `json:"fullName"`
	Link               string  `json:"link"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	PrimaryNumber      string  `json:"primaryNumber"`
	BirthDate          string  `json:"birthDate"`
	BirthStateProvince *string `json:"birthStateProvince"`
	CurrentAge         int     `json:"currentAge"`
	BirthCity          string  `json:"birthCity"`
	BirthCountry       string  `json:"birthCountry"`
	Nationality        string  `json:"nationality"`
	Height             string  `json:"height"`
	Weight             int     `json:"weight"`
	Active             bool    `json:"active"`
	AlternateCaptain   bool    `json:"alternateCaptain"`
	Captain            bool    `json:"captain"`
	Rookie             bool    `json:"rookie"`
	ShootsCatches      string  `json:"shootsCatches"`
	RosterStatus       string  `json:"rosterStatus"`
	CurrentTeam        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"currentTeam"`
	PrimaryPosition struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"primaryPosition"`
}

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		TimeZone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division        struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		NameShort    string `json:"nameShort"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId"`
		TeamName    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	ShortName       string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}

type League struct {
	ID   *int   `json:"id"`
	Name string `json:"name"`
}

type Division struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	Conference   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Active bool `json:"active"`
}

type SeasonStat struct {
	Season string `json:"season"`
	Team   Team   `json:"team"`
	League League `json:"league"`
}

type ForwardSeasonStat struct {
	SeasonStat
	Stat ForwardStat `json:"stat"`
}

// ForwardStats is the stats for a forward
type ForwardStat struct {
	Assists              int      `json:"assists"`
	Goals                int      `json:"goals"`
	Pim                  int      `json:"pim"`
	Games                int      `json:"games"`
	PowerPlayGoals       int      `json:"powerPlayGoals"`
	PenaltyMinutes       string   `json:"penaltyMinutes"`
	ShortHandedGoals     int      `json:"shortHandedGoals"`
	PlusMinus            int      `json:"plusMinus"`
	Points               int      `json:"points"`
	TimeOnIce            *string  `json:"timeOnIce"`
	Shots                *int     `json:"shots"`
	Hits                 *int     `json:"hits"`
	PowerPlayPoints      *int     `json:"powerPlayPoints"`
	PowerPlayTimeOnIce   *string  `json:"powerPlayTimeOnIce"`
	EvenTimeOnIce        *string  `json:"evenTimeOnIce"`
	FaceOffPct           *float64 `json:"faceOffPct"`
	ShotPct              *float64 `json:"shotPct"`
	GameWinningGoals     *int     `json:"gameWinningGoals"`
	OverTimeGoals        *int     `json:"overTimeGoals"`
	ShortHandedPoints    *int     `json:"shortHandedPoints"`
	ShortHandedTimeOnIce *string  `json:"shortHandedTimeOnIce"`
	Blocked              *int     `json:"blocked"`
	Shifts               *int     `json:"shifts"`
}

type GoalieSeasonStat struct {
	SeasonStat
	Stat GoalieStat `json:"stat"`
}

// GoalieStats is the stats for a goalie
type GoalieStat struct {
	TimeOnIce                  string   `json:"timeOnIce"`
	Shutouts                   int      `json:"shutouts"`
	Wins                       int      `json:"wins"`
	Losses                     int      `json:"losses"`
	Saves                      int      `json:"saves"`
	SavePercentage             float64  `json:"savePercentage"`
	GoalAgainstAverage         float64  `json:"goalAgainstAverage"`
	Games                      int      `json:"games"`
	ShotsAgainst               int      `json:"shotsAgainst"`
	GoalsAgainst               int      `json:"goalsAgainst"`
	Ties                       *int     `json:"ties"`
	Ot                         *int     `json:"ot"`
	PowerPlaySaves             *int     `json:"powerPlaySaves"`
	ShortHandedSaves           *int     `json:"shortHandedSaves"`
	EvenSaves                  *int     `json:"evenSaves"`
	ShortHandedShots           *int     `json:"shortHandedShots"`
	EvenShots                  *int     `json:"evenShots"`
	PowerPlayShots             *int     `json:"powerPlayShots"`
	GamesStarted               *int     `json:"gamesStarted"`
	PowerPlaySavePercentage    *float64 `json:"powerPlaySavePercentage"`
	ShortHandedSavePercentage  *float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage *float64 `json:"evenStrengthSavePercentage"`
}

type Season struct {
	SeasonID               string `json:"seasonId"`
	RegularSeasonStartDate string `json:"regularSeasonStartDate"`
	RegularSeasonEndDate   string `json:"regularSeasonEndDate"`
	SeasonEndDate          string `json:"seasonEndDate"`
	NumberOfGames          int    `json:"numberOfGames"`
	TiesInUse              bool   `json:"tiesInUse"`
	OlympicsParticipation  bool   `json:"olympicsParticipation"`
	ConferencesInUse       bool   `json:"conferencesInUse"`
	DivisionsInUse         bool   `json:"divisionsInUse"`
	WildCardInUse          bool   `json:"wildCardInUse"`
}
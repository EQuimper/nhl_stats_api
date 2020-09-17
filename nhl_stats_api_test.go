package nhl_stats_api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NHLStatsAPISuite struct {
	suite.Suite
	api        *Client
	testServer *httptest.Server
}

func (suite *NHLStatsAPISuite) SetupTest() {
	mux := http.NewServeMux()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetTeams.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetTeams.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetTeams.json: %v", err)
		}
	})
	mux.HandleFunc("/divisions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetDivisions.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetDivisions.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetDivisions.json: %v", err)
		}
	})
	mux.HandleFunc("/conferences", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetConferences.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetConferences.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetConferences.json: %v", err)
		}
	})
	mux.HandleFunc("/positions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetPositions.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetPositions.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetPositions.json: %v", err)
		}
	})
	mux.HandleFunc("/seasons", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetSeasons.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetSeasons.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetSeasons.json: %v", err)
		}
	})
	mux.HandleFunc("/seasons/current", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		file, err := getJsonFile("test_data/GetSeasonsCurrent.json")
		if err != nil {
			suite.T().Fatalf("error while opening GetSeasonsCurrent.json: %v", err)
		}

		_, err = w.Write(file)
		if err != nil {
			suite.T().Fatalf("error while writing GetSeasonsCurrent.json: %v", err)
		}
	})

	suite.testServer = httptest.NewServer(mux)

	suite.api = NewClient(&suite.testServer.URL)
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetTeams() {
	res, err := suite.api.GetTeams()
	suite.Require().NoError(err)

	suite.Require().Len(res, 31)

	for _, team := range res {
		suite.Require().NotEmpty(team.ID, "team.ID is empty %v", team.ID)
		suite.Require().NotEmpty(team.Name, "team.Name is empty %v", team.Name)
	}
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetDivisions() {
	res, err := suite.api.GetDivisions()
	suite.Require().NoError(err)

	suite.Require().Len(res, 4)

	for _, division := range res {
		suite.Require().NotEmpty(division.ID, "division.ID is empty %v", division.ID)
		suite.Require().NotEmpty(division.Name, "division.Name is empty %v", division.Name)
	}
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetConferences() {
	res, err := suite.api.GetConferences()
	suite.Require().NoError(err)

	suite.Require().Len(res, 2)

	for _, conference := range res {
		suite.Require().NotEmpty(conference.ID, "conference.ID is empty %v", conference.ID)
		suite.Require().NotEmpty(conference.Name, "conference.Name is empty %v", conference.Name)
	}
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetSeasons() {
	res, err := suite.api.GetSeasons()
	suite.Require().NoError(err)

	suite.Require().Len(res, 102)

	for _, season := range res {
		suite.Require().NotEmpty(season.SeasonID, "season.SeasonID is empty %v", season.SeasonID)
		suite.Require().NotEmpty(season.RegularSeasonStartDate, "season.RegularSeasonStartDate is empty %v", season.RegularSeasonStartDate)
	}
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetCurrentSeason() {
	season, err := suite.api.GetCurrentSeason()
	suite.Require().NoError(err)

	suite.Require().NotEmpty(season.SeasonID, "season.SeasonID is empty %v", season.SeasonID)
	suite.Require().NotEmpty(season.RegularSeasonStartDate, "season.RegularSeasonStartDate is empty %v", season.RegularSeasonStartDate)
}

func (suite *NHLStatsAPISuite) TestStatsApi_GetPositions() {
	res, err := suite.api.GetPositions()
	suite.Require().NoError(err)

	suite.Require().Len(res, 7)

	for _, position := range res {
		suite.Require().NotEmpty(position.Code, "position.Code is empty %v", position.Code)
		suite.Require().NotEmpty(position.Type, "position.Type is empty %v", position.Type)
		suite.Require().NotEmpty(position.FullName, "position.FullName is empty %v", position.FullName)
		suite.Require().NotEmpty(position.Abbrev, "position.Abbrev is empty %v", position.Abbrev)
	}
}

func TestNHLStatsApiSuite(t *testing.T) {
	suite.Run(t, new(NHLStatsAPISuite))
}

func (suite *NHLStatsAPISuite) TearDownTest() {
	defer suite.testServer.Close()
}

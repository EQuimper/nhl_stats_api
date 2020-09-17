package nhl_stats_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
}

func NewClient(baseUrl *string) *Client {
	var url string

	if baseUrl == nil {
		url = "https://statsapi.web.nhl.com/api/v1"
	} else {
		url = *baseUrl
	}

	return &Client{
		BaseUrl: url,
	}
}

// GetTeams return all the current nhl teams
func (s *Client) GetTeams() ([]Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams", s.BaseUrl))
	if err != nil {
		return nil, fmt.Errorf("GetTeams http error: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		Teams []Team `json:"teams"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Teams, err
}

// GetPlayerInfo return the player info from the given id
func (s *Client) GetPlayerInfo(playerID int) (PlayerInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/people/%d", s.BaseUrl, playerID))
	if err != nil {
		return PlayerInfo{}, fmt.Errorf("GetPlayerInfo http error: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		People []PlayerInfo `json:"people"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.People[0], err
}

// GetDivisions return all the nhl division
func (s *Client) GetDivisions() ([]Division, error) {
	res, err := http.Get(fmt.Sprintf("%s/divisions", s.BaseUrl))
	if err != nil {
		return nil, fmt.Errorf("GetDivisions http error: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		Divisions []Division `json:"divisions"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Divisions, err
}

// GetConferences return all the nhl conferences
func (s *Client) GetConferences() ([]Conference, error) {
	res, err := http.Get(fmt.Sprintf("%s/conferences", s.BaseUrl))
	if err != nil {
		return nil, fmt.Errorf("GetConferences http error: %v", err)
	}
	defer res.Body.Close()

	var response struct {
		Conferences []Conference `json:"conferences"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Conferences, err
}

// GetPositions return all the positions
func (s *Client) GetPositions() ([]Position, error) {
	res, err := http.Get(fmt.Sprintf("%s/positions", s.BaseUrl))
	if err != nil {
		return nil, fmt.Errorf("GetPositions http error: %v", err)
	}
	defer res.Body.Close()

	var response []Position

	err = json.NewDecoder(res.Body).Decode(&response)

	return response, err
}

// GetSeasons return all the seasons
func (s *Client) GetSeasons() ([]Season, error) {
	res, err := http.Get(fmt.Sprintf("%s/seasons", s.BaseUrl))
	if err != nil {
		return nil, fmt.Errorf("GetSeasons http error: %v", err)
	}
	defer res.Body.Close()

	var response struct{
		Seasons []Season `json:"seasons"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Seasons, err
}

// GetCurrentSeason return the current season
func (s *Client) GetCurrentSeason() (Season, error) {
	res, err := http.Get(fmt.Sprintf("%s/seasons/current", s.BaseUrl))
	if err != nil {
		return Season{}, fmt.Errorf("GetCurrentSeason http error: %v", err)
	}
	defer res.Body.Close()

	var response struct{
		Seasons []Season `json:"seasons"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Seasons[0], err
}


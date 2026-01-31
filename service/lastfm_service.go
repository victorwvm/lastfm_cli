package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victorwvm/lastfm_cli/model"
)

const baseURL = "http://ws.audioscrobbler.com/2.0/"

func GetTopTracks(apiKey string, user string) (*model.TopTrackResponse, error) {

	url := fmt.Sprintf("%s?method=user.gettoptracks&user=%s&api_key=%s&format=json", baseURL, user, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.TopTrackResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetTopArtist(apiKey string, user string) (*model.TopArtistResponse, error) {

	url := fmt.Sprintf("%s?method=user.gettopartists&user=%s&api_key=%s&format=json", baseURL, user, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.TopArtistResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func GetTopRecents(apiKey, user string) (*model.TopArtistResponse, error) {

	url := fmt.Sprintf("%s?method=user.gettopartists&user=%s&api_key=%s&format=json", baseURL, user, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result model.TopArtistResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

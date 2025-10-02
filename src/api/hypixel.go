package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/utility"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	jsoniter "github.com/json-iterator/go"
)

var HYPIXEL_API_KEY = os.Getenv("HYPIXEL_API_KEY")

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

<<<<<<< HEAD
func GetPlayer(uuid string) (*models.Player, error) {
	rawReponse := &models.HypixelPlayerResponse{}
	response := &models.Player{}
=======
func GetPlayer(uuid string) (*skycrypttypes.Player, error) {
	var rawReponse models.HypixelPlayerResponse
	var response skycrypttypes.Player
>>>>>>> 51b9d29d5b4157f70b06d9ae12a7d3e9003cf644

	if !utility.IsUUID(uuid) {
		respUUID, err := GetUUID(uuid)
		if err != nil {
			return response, err
		}

		uuid = respUUID
	}

	cache, err := redis.Get(fmt.Sprintf(`player:%s`, uuid))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), rawReponse)
		if err == nil {
			return rawReponse.Player, nil
		}
	}

	resp, err := httpClient.Get(fmt.Sprintf("https://api.hypixel.net/v2/player?key=%s&uuid=%s", HYPIXEL_API_KEY, uuid))

	if err != nil {
		return response, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return response, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, rawReponse)
	if err != nil {
		return rawReponse.Player, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`player:%s`, uuid), string(body), 24*60*60)
	return rawReponse.Player, nil
}

func GetProfiles(uuid string) (*models.HypixelProfilesResponse, error) {
	response := &models.HypixelProfilesResponse{}
	if !utility.IsUUID(uuid) {
		respUUID, err := GetUUID(uuid)
		if err != nil {
			return response, err
		}

		uuid = respUUID
	}

	cache, err := redis.Get(fmt.Sprintf(`profiles:%s`, uuid))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), response)
		if err == nil {
			return response, nil
		}
	}

	resp, err := httpClient.Get(fmt.Sprintf("https://api.hypixel.net/v2/skyblock/profiles?key=%s&uuid=%s", HYPIXEL_API_KEY, uuid))
	if err != nil {
		return response, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return response, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return response, fmt.Errorf("error parsing JSON: %v", err)
	}

	if response.Cause != "" && !response.Success {
		return response, fmt.Errorf("error fetching profiles: %s", response.Cause)
	}

	redis.Set(fmt.Sprintf(`profiles:%s`, uuid), string(body), 5*60) // Cache for 5 minutes
	return response, nil
}

func GetProfile(uuid string, profileId ...string) (*skycrypttypes.Profile, error) {
	profiles, err := GetProfiles(uuid)
	if err != nil {
		return &skycrypttypes.Profile{}, err
	}

	// If no profileId provided, return the first profile or selected profile
	if len(profileId) == 0 || (len(profileId) == 1 && profileId[0] == "") {
		if len(profiles.Profiles) == 0 {
			return &skycrypttypes.Profile{}, fmt.Errorf("no profiles found for UUID %s", uuid)
		}

		for _, profile := range profiles.Profiles {
			if profile.Selected {
				return &profile, nil
			}
		}

		return &profiles.Profiles[0], nil
	}

	// If profileId is provided, search for it
	targetProfileId := profileId[0]
	for _, profile := range profiles.Profiles {
		if profile.ProfileID == targetProfileId || profile.CuteName == targetProfileId {
			return &profile, nil
		}
	}

	return &skycrypttypes.Profile{}, fmt.Errorf("profile with ID %s not found for UUID %s", targetProfileId, uuid)
}

<<<<<<< HEAD
func GetMuseum(profileId string) (map[string]*models.Museum, error) {
	rawReponse := &models.HypixelMuseumResponse{}
=======
func GetMuseum(profileId string) (map[string]*skycrypttypes.Museum, error) {
	var rawReponse models.HypixelMuseumResponse
>>>>>>> 51b9d29d5b4157f70b06d9ae12a7d3e9003cf644

	cache, err := redis.Get(fmt.Sprintf(`museum:%s`, profileId))
	if err == nil && cache != "" {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal([]byte(cache), rawReponse)
		if err == nil {
			return rawReponse.Members, nil
		}
	}

	resp, err := httpClient.Get(fmt.Sprintf("https://api.hypixel.net/v2/skyblock/museum?key=%s&profile=%s", HYPIXEL_API_KEY, profileId))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, rawReponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`museum:%s`, profileId), string(body), 60*30) // Cache for 30 minutes
	return rawReponse.Members, nil
}

<<<<<<< HEAD
func GetGarden(profileId string) (*models.GardenRaw, error) {
	rawReponse := &models.HypixelGardenResponse{}
=======
func GetGarden(profileId string) (*skycrypttypes.Garden, error) {
	var rawReponse models.HypixelGardenResponse
>>>>>>> 51b9d29d5b4157f70b06d9ae12a7d3e9003cf644

	cache, err := redis.Get(fmt.Sprintf(`garden:%s`, profileId))
	if err == nil && cache != "" {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal([]byte(cache), rawReponse)
		if err == nil {
			return rawReponse.Garden, nil
		}
	}

	resp, err := httpClient.Get(fmt.Sprintf("https://api.hypixel.net/v2/skyblock/garden?key=%s&profile=%s", HYPIXEL_API_KEY, profileId))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, rawReponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf(`garden:%s`, profileId), string(body), 60*30) // Cache for 30 minutes
	return rawReponse.Garden, nil
}

package api

import (
	"encoding/json"
	"fmt"
	"io"
	redis "skycrypt/src/db"
	"skycrypt/src/models"
	"skycrypt/src/utility"

	"net/http"
	"strings"
)

func GetUUID(username string) (string, error) {
	var post models.MowojangReponse

	cache, err := redis.Get(fmt.Sprintf("uuid:%s", strings.ToLower(username)))
	if err == nil && cache != "" {
		return cache, nil
	}

	resp, err := http.Get(fmt.Sprintf("https://mowojang.matdoes.dev/%s", username))
	if err != nil {
		return post.UUID, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return post.UUID, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		return post.UUID, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf("uuid:%s", strings.ToLower(post.Name)), post.UUID, 24*60*60) // Cache for 24 hours
	redis.Set(fmt.Sprintf("username:%s", post.UUID), post.Name, 24*60*60)              // Cache for 24 hours

	return post.UUID, nil
}

func GetUsername(uuid string) (string, error) {
	var post models.MowojangReponse

	cache, err := redis.Get(fmt.Sprintf("username:%s", uuid))
	if err == nil && cache != "" {
		return cache, nil
	}

	resp, err := http.Get(fmt.Sprintf("https://mowojang.matdoes.dev/%s", uuid))
	if err != nil {
		return post.Name, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return post.Name, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		return post.Name, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf("uuid:%s", strings.ToLower(post.Name)), uuid, 24*60*60) // Cache for 24 hours
	redis.Set(fmt.Sprintf("username:%s", uuid), post.Name, 24*60*60)              // Cache for 24 hours

	return post.Name, nil
}

func ResolvePlayer(uuid string) (*models.MowojangReponse, error) {
	var post models.MowojangReponse
	if !utility.IsUUID(uuid) {
		tempUUID, err := GetUUID(uuid)
		if err != nil {
			return &post, fmt.Errorf("error resolving UUID for username '%s': %v", uuid, err)
		}
		uuid = tempUUID
	}

	cache, err := redis.Get(fmt.Sprintf("mowojang:%s", uuid))
	if err == nil && cache != "" {
		err = json.Unmarshal([]byte(cache), &post)
		if err == nil {
			return &post, nil
		}
	}

	resp, err := http.Get(fmt.Sprintf("https://mowojang.matdoes.dev/%s", uuid))
	if err != nil {
		return &post, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &post, fmt.Errorf("error reading response: %v", err)
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		return &post, fmt.Errorf("error parsing JSON: %v", err)
	}

	redis.Set(fmt.Sprintf("mowojang:%s", uuid), string(body), 24*60*60) // Cache for 24 hours

	return &post, nil
}

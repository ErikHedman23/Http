package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	marshalledUser, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(marshalledUser))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	var updatedUser User
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&updatedUser); err != nil {
		return User{}, err
	}

	return updatedUser, nil

}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()
	var retrievedUser User
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&retrievedUser); err != nil {
		return User{}, nil
	}

	return retrievedUser, nil

}

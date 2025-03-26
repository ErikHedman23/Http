package main

import (
	"errors"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id
	var errr = errors.New("Error when trying to delete the user: " + string(id))

	req, err := http.NewRequest("DELETE", fullURL, nil)

	if err != nil {
		return errr
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return errr
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errr
	}

	return nil
}

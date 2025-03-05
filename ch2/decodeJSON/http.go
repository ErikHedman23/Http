package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	for _, issue := range issues {
		fmt.Printf("Issue - Title: %v, Estimate: %v", issue.Title, issue.Estimate)
	}
	return issues, nil
	// ?
}

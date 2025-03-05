package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading request: %w", err)
	}

	return resp, nil
}

package main

import (
	"net/url"
	"strings"
)

func normalizeURL(baseURL string) (string, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", err

	}

	outputURL := parsedURL.Host + strings.TrimSuffix(parsedURL.Path, "/")

	return outputURL, nil
}

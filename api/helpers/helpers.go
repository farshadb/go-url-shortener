package helpers

import (
	"os"
	"strings"
)

func EnforceHTTPS () string {
	if url[:4] != "https" {
		return "https://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := string.Replace(url, "http://", "", 1)
	newURL  = string.Replace(newURL, "https://", "", 1) 
	newURL  = string.Replace(newURL, "www.", "", 1)
	newURL  = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
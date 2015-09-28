package plex

import (
	"log"
	"net/http"
)

const (
	PLEX_REFRESH_URL = "http://127.0.0.1:32400/library/sections/1/refresh"
)

func Refresh() bool {
	_, err := http.Get(PLEX_REFRESH_URL)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

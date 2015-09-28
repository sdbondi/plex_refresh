package plex

import (
	"fmt"
	"log"
	"net/http"
)

const (
	PLEX_REFRESH_URL = "http://127.0.0.1:32400/library/sections/%d/refresh"
)

func Refresh(sectionId int) bool {
	_, err := http.Get(fmt.Sprintf(PLEX_REFRESH_URL, sectionId))
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

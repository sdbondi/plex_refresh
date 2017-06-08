package plex

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	PLEX_REFRESH_URL = "http://127.0.0.1:32400/library/sections/%d/refresh?X-Plex-Token=%s"
)

func Refresh(token string, sectionId int) bool {
	fmt.Println("Triggering refresh!")
	url := fmt.Sprintf(PLEX_REFRESH_URL, sectionId, token)
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			log.Printf("Refresh succeeded")
		} else {
			log.Printf("Error: Invalid response code %d", resp.StatusCode)

			data, ioErr := ioutil.ReadAll(resp.Body)
			if ioErr != nil {
				log.Fatal(ioErr)
				return false
			}

			if len(data) > 0 {
				fmt.Println("Response: " + string(data))
			}
		}
	}

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

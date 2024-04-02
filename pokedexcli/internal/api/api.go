package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetLocationsData(url *string) (LocationsData, error) {
	res, err := http.Get(*url)
	if err != nil {
		// log.Fatal("Error accessing locations api: ", err)
		log.Println("Error accessing locations api: ", err)
		return LocationsData{}, err
	}
	defer res.Body.Close()

	var locationsData LocationsData
	if err := json.NewDecoder(res.Body).Decode(&locationsData); err != nil {
		log.Println("Error decoding JSON: ", err)
		return LocationsData{}, err
	}

	return locationsData, nil
}

type LocationsData struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

package locations

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetLocationsData(url *string) (LocationsData, []byte, error) {
	res, err := http.Get(*url)
	if err != nil {
		log.Println("Error accessing locations api: ", err)
		return LocationsData{}, nil, err
	}

	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsData{}, nil, err
	}

	var locationsData LocationsData
	err = json.Unmarshal(dat, &locationsData)
	if err != nil {
		return LocationsData{}, nil, err
	}

	return locationsData, dat, nil
}

type LocationsData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

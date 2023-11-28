package groupie_tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ReadArtistApi() ([]Artists, error) {

	var artists []Artists

	r, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, errors.New("error get artists api request")
	}
	defer r.Body.Close()
	data, resp_err := io.ReadAll(r.Body)
	if resp_err != nil {
		return nil, errors.New("error reading json artists")
	}
	marsh_err := json.Unmarshal(data, &artists)
	if marsh_err != nil {
		return nil, errors.New("error during unmarshal artists data")
	}
	return artists, nil
}

func ReadLocationApi(url string) (Locations, error) {
	locations := Locations{}

	r, err := http.Get(url)
	if err != nil {
		return locations, errors.New("error get location api request")
	}
	defer r.Body.Close()
	data, resp_err := io.ReadAll(r.Body)
	if resp_err != nil {
		return locations, errors.New("error reading json location")
	}
	marsh_err := json.Unmarshal(data, &locations)
	if marsh_err != nil {
		return locations, errors.New("error during unmarshal location data")
	}
	return locations, nil
}

func ReadDatesApi(url string) (Dates, error) {
	dates := Dates{}

	r, err := http.Get(url)
	if err != nil {
		return dates, errors.New("error get dates api request")
	}
	defer r.Body.Close()
	data, resp_err := io.ReadAll(r.Body)
	if resp_err != nil {
		return dates, errors.New("error reading json dates")
	}
	marsh_err := json.Unmarshal(data, &dates)
	if marsh_err != nil {
		return dates, errors.New("error during unmarshal dates data")
	}
	return dates, nil
}

func ReadRelationsApi(url string) (Relations, error) {
	relation := Relations{}

	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error read relations url")
	}
	defer r.Body.Close()
	data, resp_err := io.ReadAll(r.Body)
	if resp_err != nil {
		fmt.Println("Error read body data relations")
	}
	marsh_err := json.Unmarshal(data, &relation)
	if marsh_err != nil {
		fmt.Println("error unmarshal relations")
	}
	return relation, nil
}

func PopulateDetails(artists Artists, locations Locations, dates Dates, relation Relations, geocodeout []GeoCodeOut) Full {
	return Full{Artists: artists, Locations: locations, Dates: dates, Relations: relation, GeoCodeOut: geocodeout}
}

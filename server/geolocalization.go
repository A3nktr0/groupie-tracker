package groupie_tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Read geocode api and keep latitude and longitude of location
func ReadGeoAPI(ch chan []GeoCodeOut, locations []string) {
	var geoCode []GeoCode
	var position GeoCodeOut
	var geoCodeOut []GeoCodeOut

	for i := range locations {
		tmp := strings.ReplaceAll(locations[i], "-", "%2C")

		r, err := http.Get(fmt.Sprintf("https://nominatim.openstreetmap.org/search.php?q=%s&accept-language=en&limit=1&format=jsonv2", tmp))
		if err != nil {
			fmt.Println("Error read geoCode url")
		}
		defer r.Body.Close()
		data, resp_err := io.ReadAll(r.Body)
		if resp_err != nil {
			fmt.Println("Error read body data geoCode")
		}

		marsh_err := json.Unmarshal(data, &geoCode)
		if marsh_err != nil {
			fmt.Println("error unmarshal geoCode")
		}

		for i := range geoCode {
			position.Name = geoCode[i].Name
			position.Latitude, _ = strconv.ParseFloat(geoCode[i].Latitude, 32)
			position.Longitude, _ = strconv.ParseFloat(geoCode[i].Longitude, 32)
			geoCodeOut = append(geoCodeOut, position)
		}
	}

	ch <- geoCodeOut
}

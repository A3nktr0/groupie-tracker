package groupie_tracker

// type Data interface {
// 	Filters(res chan<- []Artists, data []Artists, query string)
// }

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type GeoCode struct {
	Name      string `json:"display_name"`
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
}
type GeoCodeOut struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type Full struct {
	Artists    Artists
	Locations  Locations
	Dates      Dates
	Relations  Relations
	GeoCodeOut []GeoCodeOut
}

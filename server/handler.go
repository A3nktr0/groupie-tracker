package groupie_tracker

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Return Error status code if an error occurred
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status) // Write status code
	switch status {
	case http.StatusNotFound:
		tmpl := template.Must(template.ParseFiles("templates/404.html", "templates/header.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
		}
	case http.StatusBadRequest:
		fmt.Fprint(w, http.StatusText(status))
	case http.StatusInternalServerError:
		fmt.Fprint(w, http.StatusText(status))
	}
}

// Home handle function, serve template and and display homepage
func home(w http.ResponseWriter, r *http.Request) {
	var creationDateMin string
	var creationDateMax string
	var firstAlbumMin string
	var firstAlbumMax string
	var membersNbr []string
	var concertLoc string
	var resultSearch []Full
	var tmp Full

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/header.html")) // Define template

	if r.Method == "GET" {

		rawData, e := ReadArtistApi()
		if e != nil {
			errorHandler(w, r, http.StatusInternalServerError)
		}

		//Parse form when method are GET and define variables for filtering
		creationDateMin = r.FormValue("slider-left")
		creationDateMax = r.FormValue("slider-right")
		firstAlbumMin = r.FormValue("slider-left-album")
		firstAlbumMax = r.FormValue("slider-right-album")
		concertLoc = strings.ToLower(strings.ReplaceAll(r.FormValue("location"), ", ", "-"))

		// Verify if multiple checkboxes are checked
		for i := 0; i <= 7; i++ {
			if r.URL.Query().Has(fmt.Sprintf("m%d", i)) {
				membersNbr = append(membersNbr, r.FormValue(fmt.Sprintf("m%d", i)))
			}
		}

		// Apply Filter
		data := FilteredData(rawData, creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax, membersNbr, concertLoc)

		if r.URL.Query().Has("search") {

			tmpArr := SearchBar(r.URL.Query().Get("search"), rawData)
			resultSearch = append(resultSearch, tmpArr...)
			err := tmpl.Execute(w, resultSearch)
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
		} else {
			for i := range data {
				tmp.Artists = data[i]
				tmp.Locations, _ = ReadLocationApi(data[i].Locations)
				resultSearch = append(resultSearch, tmp)
			}
			err := tmpl.Execute(w, resultSearch)
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
		}
	} else {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
}

// Artists handle function, serve template and and display artist page
func artists(w http.ResponseWriter, r *http.Request) {
	var full_data Full
	var resultSearch []Full

	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		artists, e := ReadArtistApi()
		if e != nil {
			errorHandler(w, r, http.StatusInternalServerError)
		}

		if r.URL.Query().Has("search") {
			tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/header.html")) // Define template

			tmpArr := SearchBar(r.URL.Query().Get("search"), artists)
			resultSearch = append(resultSearch, tmpArr...)
			err := tmpl.Execute(w, resultSearch)
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
		} else {

			tmp, e := strconv.Atoi(id)
			if e != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
			if tmp < 0 || tmp > len(artists) {
				errorHandler(w, r, http.StatusNotFound)
				return
			}
			locations, e := ReadLocationApi(artists[tmp-1].Locations)
			if e != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}

			tmp_geo_code := make(chan []GeoCodeOut)
			go ReadGeoAPI(tmp_geo_code, locations.Locations)
			geocode := <-tmp_geo_code

			dates, e := ReadDatesApi(artists[tmp-1].ConcertDates)
			if e != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
			relations, e := ReadRelationsApi(artists[tmp-1].Relations)
			if e != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}

			full_data = PopulateDetails(artists[tmp-1], locations, dates, relations, geocode)

			resultSearch = append(resultSearch, full_data)

			tmpl := template.Must(template.ParseFiles("templates/artists.html", "templates/header.html"))
			err := tmpl.Execute(w, resultSearch[0])
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
			}
		}

	} else {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

}

func Handlers() {

	http.HandleFunc("/", home)
	http.HandleFunc("/artists", artists)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer((http.Dir("static")))))

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

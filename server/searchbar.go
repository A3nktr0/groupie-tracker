package groupie_tracker

import (
	"encoding/json"
	"strconv"
	"strings"
)

func SearchBar(query string, data []Artists) []Full {
	var output []Full
	var tmp Full

	str := strings.Split(query, " (")

	for i := range data {
		tmp.Artists = data[i]
		id := data[i].ID
		tmp.Locations, _ = ReadLocationApi(data[id-1].Locations)
		out, _ := json.Marshal(tmp)

		if strings.Contains(string(out), str[0]) || strings.Contains(string(out), strings.Title(str[0])) {
			date, _ := strconv.Atoi(str[0])

			switch {
			case strings.ToTitle(str[0]) == tmp.Artists.Name:
				output = append(output, tmp)
			case strings.Contains(strings.Join(tmp.Artists.Members, " "), str[0]):
				output = append(output, tmp)
			case date == tmp.Artists.CreationDate:
				output = append(output, tmp)
			case str[0] == tmp.Artists.FirstAlbum:
				output = append(output, tmp)
			case strings.Contains(strings.Join(tmp.Locations.Locations, " "), str[0]):
				output = append(output, tmp)
			default:
				output = append(output, tmp)
			}
		}
	}
	return output
}

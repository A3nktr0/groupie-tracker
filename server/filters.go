package groupie_tracker

import (
	"strconv"
	"strings"
)

// Apply filter creation date
func FilterCreationDate(res chan []Artists, data []Artists, query1, query2 string) {
	var out []Artists
	filterMin, _ := strconv.Atoi(query1)
	filterMax, _ := strconv.Atoi(query2)

	for i := range data {
		if data[i].CreationDate >= filterMin && data[i].CreationDate <= filterMax {
			out = append(out, data[i])
		}
	}
	res <- out
}

// Apply filter first album date
func FilterFirstAlbum(res chan []Artists, data []Artists, query1, query2 string) {
	var out []Artists
	for i := range data {
		tmpData := strings.Split(data[i].FirstAlbum, "-")
		if tmpData[2] >= query1 && tmpData[2] <= query2 {
			out = append(out, data[i])
		}
	}
	res <- out
}

// Apply filter number of members
func FilterMembers(res chan []Artists, data []Artists, query []string) {
	var out []Artists
	for i := range query {
		filter, _ := strconv.Atoi(query[i])
		for j := range data {
			if len(data[j].Members) == filter {
				out = append(out, data[j])
			}
		}
	}
	res <- out
}

// Apply filter concerts location
func FilteredLocation(res chan []Artists, data []Artists, query string) {
	var out []Artists
	for i := range data {
		tmp, _ := ReadLocationApi(data[i].Locations)
		for j := range tmp.Locations {
			if tmp.Locations[j] == query {
				out = append(out, data[i])
			}
		}
	}
	res <- out
}

// Compare first and second artists arrays
func CompareAB(ch chan []Artists, a, b []Artists) {
	var ab []Artists
	for i := range a {
		for j := range b {
			if a[i].ID == b[j].ID {
				ab = append(ab, a[i])
			}
		}
	}
	ch <- ab
}

// Compare first and third artists arrays
func CompareAC(ch chan []Artists, a, c []Artists) {
	var ac []Artists
	for i := range a {
		for j := range c {
			if a[i].ID == c[j].ID {
				ac = append(ac, a[i])
			}
		}
	}
	ch <- ac
}

// Compare first and fourth artists arrays
func CompareAD(ch chan []Artists, a, d []Artists) {
	var ad []Artists
	for i := range a {
		for j := range d {
			if a[i].ID == d[j].ID {
				ad = append(ad, a[i])
			}
		}
	}
	ch <- ad
}

// Main filter
func FilteredData(rawData []Artists, creationDateMin, creationDateMax, firstAlbumMin, firstAlbumMax string, membersNbr []string, concertLoc string) []Artists {
	var filteredC, filteredFA, filteredM, filteredL, filtered, filteredOut []Artists

	if creationDateMin == "" && creationDateMax == "" {
		filteredC = append(filteredC, rawData...)
	} else {
		date := make(chan []Artists)
		go FilterCreationDate(date, rawData, creationDateMin, creationDateMax)
		filteredC = <-date
	}

	if firstAlbumMin == "" && firstAlbumMax == "" {
		filteredFA = append(filteredFA, rawData...)
	} else {
		firstAlbum := make(chan []Artists)
		go FilterFirstAlbum(firstAlbum, rawData, firstAlbumMin, firstAlbumMax)
		filteredFA = <-firstAlbum
	}

	if len(membersNbr) == 0 || membersNbr[0] == "0" {
		filteredM = append(filteredM, rawData...)
	} else {
		members := make(chan []Artists)
		go FilterMembers(members, rawData, membersNbr)
		filteredM = <-members
	}

	if concertLoc == "" {
		filteredL = append(filteredL, rawData...)
	} else {
		loc := make(chan []Artists)
		go FilteredLocation(loc, rawData, concertLoc)
		filteredL = <-loc

	}

	tmp1, tmp2, tmp3 := make(chan []Artists), make(chan []Artists), make(chan []Artists)

	go CompareAB(tmp1, filteredC, filteredM)
	go CompareAC(tmp2, filteredC, filteredFA)
	go CompareAD(tmp3, filteredC, filteredL)

	compAB, compAC, compAD := <-tmp1, <-tmp2, <-tmp3

	for i := range compAB {
		for j := range compAC {
			if compAB[i].ID == compAC[j].ID {
				filtered = append(filtered, compAB[i])
			}
		}
	}
	for i := range filtered {
		for j := range compAD {
			if filtered[i].ID == compAD[j].ID {
				filteredOut = append(filteredOut, compAD[j])
			}
		}
	}
	return filteredOut
}

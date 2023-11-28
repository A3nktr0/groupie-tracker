# Groupie Tracker 

Groupie Tracker is a web application that allows users to search or to filter and view information about music artists and bands. It retrieves data from an API and provides various filtering options to refine the displayed results.

## Getting started

To run the Groupie Tracker Filters application locally, follow the steps below:

1. Start the server:

```command
    go run cmd/main.go
```

2. Access the application in your web browser at <http://localhost:8080>.

## Usage

Groupie Tracker Filters provides several filtering options to narrow down the displayed artists/bands:

- Filter by Creation Date: Use the range filter to specify a date range within which the artists/bands were formed.

- Filter by First Album Date: Use the range filter to specify a date range within which the artists/bands released their first album.

- Filter by Number of Members: Use the checkbox filter to select one or multiple options representing the desired number of members in the artists/bands.

- Filter by Locations of Concerts: enter the location in the text area.

- Enter your query in the searchbar and press enter to filter results with a written query

## Technologies Used

Go: The back-end of the application is written in Go, utilizing its concurrency features and HTTP server capabilities.

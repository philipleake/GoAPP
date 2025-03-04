package main

type MovieAPIResponse struct {
	Title  string `json:"title"`
	Year   string `json:"release_date"`
	Poster string `json:"poster_path"`
	Genre  string `json:"genre_ids"`
	ID     int    `json:"id"`
}

//TODO finish tmdb api

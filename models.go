package main

import "fmt"

type Movie struct {
	ID          int64  `bd:"id"`
	Title       string `bd:"title"`
	ReleaseDate string `bd:"releaase_date"`
	Duration    int    `bd:"duration"`
	TrailerURL  string `bd:"trailer_url"`
}

func (m Movie) String() string {
	return fmt.Sprintf("id=%v, title=%v, releaseDate=%v, duration=%v, trailer=%v",
		m.ID, m.Title, m.ReleaseDate, m.Duration, m.TrailerURL)
}

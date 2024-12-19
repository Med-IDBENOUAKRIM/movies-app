package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Runtime   Runtime   `json:"runtime"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
	Year      int32     `json:"year,omitempty"`
	CreatedAt time.Time `json:"-"`
}

// CreatedAt time.Time `json:"created_at"`

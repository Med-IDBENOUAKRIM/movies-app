package data

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/med-IDBENOUAKRIM/lets_go/internal/validator"
)

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

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "title is required")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "is required")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "is required")
	v.Check(movie.Runtime > 0, "runtime", "must be positive")

	v.Check(movie.Genres != nil, "genres", "are required")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must NOT contain duplicate values")
}

type MovieModel struct {
	DB *sql.DB
}

func (m *MovieModel) InsertMovie(movie *Movie) error {
	query := `INSERT INTO movies (title, year, runtime, genres) VALUES ($1, $2, $3, $4) RETURNING id, created_at, version`
	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
	}
	log.Println(movie)
	// ! 149
	// newMovie := Movie{
	// 	Title:   movie.Title,
	// 	Year:    movie.Year,
	// 	Runtime: movie.Runtime,
	// 	Genres:  pq.StringArray(movie.Genres),
	// }
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
	// return nil
}

func (m *MovieModel) GetMovieById(id int64) (*Movie, error) {
	query := `SELECT id, title, created_at, year, runtime, genres, version FROM movies WHERE id = $1`
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	var movie Movie
	err := m.DB.QueryRow(query, id).Scan(&movie.ID, &movie.Title, &movie.CreatedAt, &movie.Year, &movie.Runtime, pq.Array(&movie.Genres), &movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &movie, nil
}

func (m *MovieModel) UpdateMovie(movie *Movie) error {
	return nil
}

func (m *MovieModel) DeleteMovie(id int64) error {
	return nil
}

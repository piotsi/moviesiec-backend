package movies

// Movie represents details about a movie
type Movie struct {
	Title    string
	Director string
	Year     int
	Rating   float64
	Ratings  int
}

// Repository gives access to movies collection
type Repository interface {
	// GetAll gets all movies from the database
	GetAll() []Movie

	// Get gets specific movie from the database
	Get(int) (Movie, error)

	// Add adds a movie to the database
	Add(Movie) error
}

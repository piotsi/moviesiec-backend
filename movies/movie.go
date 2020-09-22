package movies

// Movie represents details about a movie
type Movie struct {
	title    string
	director string
	year     int32
	rating   float32
	ratings  int32
}

// Repository gives access to movies collection
type Repository interface {
	// GetAll gets all movies from the database
	GetAll() []Movie
}

package users

type User struct {
	StringID string `json:"stringid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Repository interface {
	SignIn() error
	SignOut() error
	SignOn() error
}

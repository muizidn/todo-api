package app

type User struct {
	uuid           string
	username       string
	email          string
	hashedPassword string
}

type UserRepo interface {
	Create(username string, email string, hashedPassword string) (*User, error)
	GetUsername(username string) (*User, error)
}

package model

type User struct {
	Email string
}

func New() User {
	return User{}
}

func (u User) getEmail() string {
	return u.Email
}


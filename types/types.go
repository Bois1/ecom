package types

import (
	"time"
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user *User) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}


func (p *RegisterUserPayload) Validate() error {
	if p.Email == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(p.Email) {
		return errors.New("invalid email format")
	}
	if p.Password == "" {
		return errors.New("password is required")
	}
	if len(p.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}
package types

import "time"


type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}

type User struct {
	ID			int 		`json:"id`
	FirstName 	string 		`json:"firstName"`
	LastName 	string 		`json:"lastName"`
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type RegiserUserPayload struct {
	FirstName 	string `json:"firstName"`
	LastName 	string `json:"lastName"`
	Email  		string `json:"email"`
	Password 	string `json:"password"`
}


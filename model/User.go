package model

import (
	"github.com/gbrlsnchs/jwt"
)

type User struct {
	Id int
	Email string
	Nickname string
	Age int
	Sex string
	Password string
}

type LoginToken struct {
	jwt.Payload
	ID int `json:"id"`
	Email string `json:"email"`
}

type UserToken struct {
	Token string
	User User
}

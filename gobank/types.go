package main

import (
	"math/rand"
	"time"
	"golang.org/x/crypto/bcrypt"
)

// the additional strings to the right, is a specification of how you set the variable names - they are caled struct field tags...
type LoginResponse struct {
	Number int64 `json:"number"`
	Token string `json:"token"`
}
type LoginRequest struct {
	Number int64 `json:"number"`
	Password string `json:"password"`
}
type TransferRequest struct {
	ToAccount int `json:"firstName"`
	LastName string `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID 		  int 		`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"lastName"`
	Number    int64 	`json:"number"`
	EncryptedPassword string `json:"string"`
	Balance   int64 	`json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, password string) (*Account,error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err 
	}
	return &Account{
		FirstName: firstName,
		LastName: lastName,
		EncryptedPassword: string(encpw),
		Number: int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}, nil 
}






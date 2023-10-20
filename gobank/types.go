package main

import (
	"math/rand"
	"time"
)

// the additional strings to the right, is a specification of how you set the variable names - they are caled struct field tags...

type CreateAccountRequest struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
}

type Account struct {
	ID 		  int 		`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"lastName"`
	Number    int64 	`json:"number"`
	Balance   int64 	`json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName: lastName,
		Number: int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}



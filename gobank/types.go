package main 

import "math/rand"


// the additional strings to the right, is a specification of how you set the variable names - they are caled struct field tags...
type Account struct {
	ID 		  int `json:"id"`
	FirstName string `json:"firsName"`
	LastName  string `json:"lastName"`
	Number    int64 `json:"number"`
	Balance   int64 `json:"balance"`
}

func NewAccount(firstName, lastname string) *Account {
	return &Account{
		ID : rand.Intn(10000),
		FirstName: firstName,
		LastName: lastname,
		Number: int64(rand.Intn(1000000)),
	}
}


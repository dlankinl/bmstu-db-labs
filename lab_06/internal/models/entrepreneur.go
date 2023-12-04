package models

import "time"

type Entrepreneur struct {
	FirstName string
	LastName  string
	Age       int
	Gender    bool
	Married   bool
	NetWorth  float32
	BirthDate time.Time
}

func NewEntrepreneur(firstName, lastName string, age int, gender, married bool, netWorth float32, birthDate time.Time) *Entrepreneur {
	return &Entrepreneur{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Gender:    gender,
		Married:   married,
		NetWorth:  netWorth,
		BirthDate: birthDate,
	}
}

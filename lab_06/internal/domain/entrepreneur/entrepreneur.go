package entrepreneur

import "time"

type Entrepreneur struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Gender    bool      `json:"gender"`
	Married   bool      `json:"married"`
	NetWorth  int       `json:"net_worth"`
	BirthDate time.Time `json:"birth_date"`
}

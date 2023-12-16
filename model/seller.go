package model

import "time"

type Seller struct {
	User
	Rating           float32
	RegistrationDate time.Time
}

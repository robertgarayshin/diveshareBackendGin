package model

import "time"

type Rent struct {
	Id        int
	RentBegin time.Time `db:"rent_begin"`
	RentEnd   time.Time `db:"rent_end"`
	Seller    int
	Renter    int
}

type UpdateRentInput struct {
	RentBegin time.Time
	RentEnd   time.Time
}

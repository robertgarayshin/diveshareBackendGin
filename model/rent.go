package model

import "time"

type Rent struct {
	RentId    int
	RentBegin time.Time
	RentEnd   time.Time
	Seller    User
	Renter    User
	Car       //TODO: добавить в бд информацию о машине
}

type UpdateRentInput struct {
	RentBegin time.Time
	RentEnd   time.Time
}

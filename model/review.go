package model

import "time"

type Review struct {
	Id             int
	From           User
	To             Seller
	ReviewDatetime time.Time
	Rating         float32
	Text           string
}

type UpdateReviewInput struct {
	Rating float32
	Text   string
}

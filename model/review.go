package model

import "time"

type Review struct {
	Id             int
	From           int
	To             int
	ReviewDatetime time.Time `db:"review_datetime"`
	Rating         float32   `db:"review_rating"`
	Text           string    `db:"review_text"`
}

type UpdateReviewInput struct {
	Rating float32
	Text   string
}

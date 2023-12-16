package repository

import (
	"diveshareBackendGin/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func (r ReviewPostgres) Create(review model.Review) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s ("from", "to", review_datetime, review_rating, review_text)
		VALUES ($1, $2, now(), $3, $4) RETURNING id`, ReviewsTable)
	row := r.db.QueryRow(query, review.From, review.To, review.Rating, review.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r ReviewPostgres) GetAll() ([]model.Review, error) {
	var reviews []model.Review
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY id`, ReviewsTable)
	if err := r.db.Select(&reviews, query); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r ReviewPostgres) GetById(reviewId int) (model.Review, error) {
	var review model.Review
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, ReviewsTable)
	if err := r.db.Get(&review, query, reviewId); err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func (r ReviewPostgres) Update(reviewId int, input model.UpdateReviewInput) error {
	query := fmt.Sprintf(`UPDATE %s SET review_rating = $1, review_text = $2 WHERE id = $3`, ReviewsTable)
	_, err := r.db.Exec(query, input.Rating, input.Text, reviewId)
	return err
}

func (r ReviewPostgres) Delete(reviewId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, ReviewsTable)
	_, err := r.db.Exec(query, reviewId)
	return err
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

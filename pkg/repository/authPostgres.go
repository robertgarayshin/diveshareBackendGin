package repository

import (
	"diveshareBackendGin/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_role, username, password_hash, name, "+
		"surname, confirmation_token, email_is_confirmed, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) "+
		"RETURNING id", UsersTable)
	row := r.db.QueryRow(query, "SELLER", user.Username, user.Password, user.Name, user.Surname, user.ConfirmationToken,
		"false", time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE username=$1", UsersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}

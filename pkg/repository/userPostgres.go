package repository

import (
	"diveshareBackendGin/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]model.User, error) {
	var users []model.User
	query := fmt.Sprintf(`SELECT email, user_role, name, surname FROM %s`, UsersTable)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetById(userId int) (model.User, error) {
	fmt.Println("Im in repo")
	var user model.User
	query := fmt.Sprintf(`SELECT username, user_role, name, surname FROM %s WHERE id = $1`,
		UsersTable)
	if err := r.db.Get(&user, query, userId); err != nil {
		return user, err
	}
	fmt.Println(user)
	return user, nil
}

func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = now() WHERE id = $1`,
		UsersTable)
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *UserPostgres) Update(userId int, input model.UpdateUserInput) error {
	return nil
}

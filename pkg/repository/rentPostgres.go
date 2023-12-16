package repository

import (
	"diveshareBackendGin/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RentPostgres struct {
	db *sqlx.DB
}

func (r RentPostgres) Create(rent model.Rent) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (rent_begin, rent_end, seller, renter)
		VALUES ($1, $2, $3, $4) RETURNING id`, RentsTable)
	row := r.db.QueryRow(query, rent.RentBegin, rent.RentEnd, rent.Seller, rent.Renter)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r RentPostgres) GetAll() ([]model.Rent, error) {
	var rents []model.Rent
	query := fmt.Sprintf(`SELECT * FROM %s ORDER BY id`, RentsTable)
	if err := r.db.Select(&rents, query); err != nil {
		return nil, err
	}

	return rents, nil
}

func (r RentPostgres) GetById(rentId int) (model.Rent, error) {
	var rent model.Rent
	query := fmt.Sprintf(`SELECT * FROM %s where id = $1`, RentsTable)
	if err := r.db.Get(&rent, query, rentId); err != nil {
		return model.Rent{}, err
	}

	return rent, nil
}

func (r RentPostgres) Update(rentId int, input model.UpdateRentInput) error {
	query := fmt.Sprintf(`UPDATE %s SET rent_begin = $1, rent_end = $2 WHERE id = $3`, RentsTable)
	_, err := r.db.Exec(query, input.RentBegin, input.RentEnd, rentId)
	return err
}

func (r RentPostgres) Delete(rentId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, RentsTable)
	_, err := r.db.Exec(query, rentId)
	return err
}

func NewRentPostgres(db *sqlx.DB) *RentPostgres {
	return &RentPostgres{db: db}
}

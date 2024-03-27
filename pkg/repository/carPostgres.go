package repository

import (
	"diveshareBackendGin/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CarPostgres struct {
	db *sqlx.DB
}

func (c CarPostgres) Create(car model.Car) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (category, brand, model, owner_id, price, produced_year,
		status, car_rating, created_at) VALUES ($1, $2, $3, $4, $5, $6, 'FREE', 0.0, now()) RETURNING id`, CarsTable)
	row := c.db.QueryRow(query, car.Category,
		car.Brand, car.Model, car.Owner, car.Price, car.Produced)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c CarPostgres) GetAll() ([]model.Car, error) {
	var cars []model.Car
	query := fmt.Sprintf(`SELECT id, category, brand, model, price, image FROM %s c ORDER BY c.id DESC`, CarsTable)
	if err := c.db.Select(&cars, query); err != nil {
		return nil, err
	}

	return cars, nil
}

func (c CarPostgres) GetById(carId int) (model.Car, error) {
	var car model.Car
	query := fmt.Sprintf(`SELECT id, category, brand, model, price, image
		 FROM %s where id = $1`, CarsTable)
	if err := c.db.Get(&car, query, carId); err != nil {
		return model.Car{}, err
	}
	return car, nil
}

func (c CarPostgres) Update(carId int, input model.UpdateCarInput) error {
	query := fmt.Sprintf(`UPDATE %s SET category = $1, brand = $2, model = $3, price = $4, produced_year = $5, 
		status = $6, updated_at = now() WHERE id = $7`,
		CarsTable)
	_, err := c.db.Exec(query, input.Category, input.Brand, input.Model, input.Price, input.Produced,
		input.Status, carId)
	return err
}

func (c CarPostgres) Delete(carId int) error {
	query := fmt.Sprintf(`UPDATE %s SET deleted_at = now() WHERE id = $1`, CarsTable)
	_, err := c.db.Exec(query, carId)
	return err
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

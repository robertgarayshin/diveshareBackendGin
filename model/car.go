package model

type Car struct {
	Id       int
	Category string
	Brand    string
	Model    string
	Owner    User
	Price    float32
	Produced string `db:"produced_year"`
	Status   string
	Rating   float32 `db:"car_rating"`
}

type UpdateCarInput struct {
	Category string
	Brand    string
	Model    string
	Price    float32
	Produced string
	Status   string
}

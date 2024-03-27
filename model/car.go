package model

type Car struct {
	Id       int
	Category string
	Brand    string
	Model    string
	Owner    int `db:"owner_id"`
	Price    string
	Produced string `db:"produced_year"`
	Status   string
	Rating   float32 `db:"car_rating"`
	Image    string
}

type UpdateCarInput struct {
	Category string
	Brand    string
	Model    string
	Price    float32
	Produced string
	Status   string
}

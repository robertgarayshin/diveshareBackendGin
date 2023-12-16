package model

type Car struct {
	Id       int
	Category string
	Brand    string
	Model    string
	Owner    Seller
	Price    float32
	Produced string
	Status   string
	Rating   float32
}

type UpdateCarInput struct {
	Category string
	Brand    string
	Model    string
	Price    float32
	Produced string
	Status   string
}

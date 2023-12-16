package model

type User struct {
	Id                int    `json:"id"`
	Username          string `json:"email"`
	Password          string `json:"password"`
	UserRole          string `json:"user_role" db:"user_role"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	ConfirmationToken string
	PasswordHash      string `db:"password_hash"`
}

type UpdateUserInput struct {
	Username string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

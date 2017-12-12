package models

type Employed struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Phone string `db:"phone" json:"phone"`
	Address string `db:"address" json:"address"`
	Photo string `db:"photo" json:"photo"`
}
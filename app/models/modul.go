package models

/**
Table Structure
Modul 
 */
type Modul struct {
	Id          int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Url    string `db:"url" json:"url"`
	Icon    string `db:"icon" json:"icon"`
	Description string `db:"description" json:"description"`
	IsDeleted   int    `db:"is_deleted" json:"is_deleted"`
}

package models

/**
Table Structure
Mini App
 */
type MiniApp struct {
	Id          int    `db:"id" json:"id"`
	CtrlName    string `db:"code" json:"code"`
	Description string `db:"description" json:"description"`
	IsDeleted   int    `db:"is_deleted" json:"is_deleted"`
}

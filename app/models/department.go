package models

type Department struct {
	Id                int    `db:"id" json:"id"`
	DepartmentName    string `db:"departmentname" json:"departmentname"`
	NumberOfEmployees string `db:"numberofemployees" json:"numberofemployees"`
}

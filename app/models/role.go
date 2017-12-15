package models

import (
	"github.com/revel/revel"
)

type Role struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	ParentRoleId string `db:"parent_role_id" json:"parent_role_id"`
	IsDeleted int `db:"is_deleted" json:"is_deleted"`
}

func (d Role) Validate(v *revel.Validation){
	v.Check(d.Name, revel.ValidRequired())
	v.Check(d.ParentRoleId, revel.ValidRequired())
}


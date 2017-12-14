package controllers

import (
	"GoglangRevelCURD/app/viewmodels/departmentvm"

	"github.com/revel/revel"
)

type DepartmentCtrl struct {
	BaseCtrl
}

func (a DepartmentCtrl) Index() revel.Result {
	vm := departmentvm.GetAllDepartmentViewModels{
		Txn: a.Txn,
	}

	model := vm.LoadAll()
	return a.Render(model)
}

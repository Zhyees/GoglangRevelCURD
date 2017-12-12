package controllers

import (
	"github.com/revel/revel"
	"GoglangRevelCURD/app/viewmodels/employedvm"
)

type EmployedCtrl struct {
	BaseCtrl
}

func (a EmployedCtrl) Index() revel.Result {
	vm := employedvm.GetAllViewModels{
		Txn: a.Txn,
	}

	model := vm.LoadAll()
	return a.Render(model)
}

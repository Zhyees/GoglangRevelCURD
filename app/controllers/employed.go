package controllers

import (
	"github.com/revel/revel"
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/viewmodels/employedvm"
	"GoglangRevelCURD/app/models"
	"GoglangRevelCURD/app/routes"
)

type EmployedCtrl struct {
	BaseCtrl
}

func (c EmployedCtrl) Index() revel.Result {
	vm := employedvm.GetAllViewModels{
		Txn: c.Txn,
	}

	model := vm.LoadAll()
	return c.Render(model)
}
func (c EmployedCtrl) AddEmployed() revel.Result {
	return c.Render()
}
func (c EmployedCtrl) AddAction(employed models.Employed) revel.Result {
	vm := employedvm.AddActionViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		Employed: employed,
	}
	model := vm.AddEmployed()
	if model.IsSuccess{
		c.Flash.Success(model.Message)
	}else {
		c.Flash.Error(model.Message)
	}
	return c.Redirect(routes.EmployedCtrl.Index())
}
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
func (c EmployedCtrl) Add() revel.Result {
	return c.Render()
}
func (c EmployedCtrl) UpdateByID(id int) revel.Result {
	vm := employedvm.GetByIdViewModels{
		Txn: c.Txn,
		Id: id,
	}

	model := vm.LoadAll()
	return c.Render(model)
}
func (c EmployedCtrl) AddAction(employed models.Employed) revel.Result {
	vm := employedvm.AddActionViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		Employed: employed,
	}
	model := vm.LoadAll()
	if model.IsSuccess{
		c.Flash.Success(model.Message)
	}else {
		c.Flash.Error(model.Message)
	}
	return c.Redirect(routes.EmployedCtrl.Index())
}
func (c EmployedCtrl) UpdateAction(id int, updateemployed models.Employed) revel.Result {
	vm := employedvm.UpdateActionViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		UpdateEmployed: updateemployed,
		Id: id,
	}
	model := vm.LoadAll()
	if model.IsSuccess{
		c.Flash.Success(model.Message)
	}else {
		c.Flash.Error(model.Message)
	}
	return c.Redirect(routes.EmployedCtrl.Index())
}
func (c EmployedCtrl) DeleteAction(id int) revel.Result{
	vm := employedvm.DeleteActionViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		Id: id,
	}

	model := vm.LoadAll()

	if model.IsSuccess == true{
		c.Flash.Success(model.Message)
	}else {
		c.Flash.Error(model.Message)
	}

	return c.Redirect(routes.EmployedCtrl.Index())
}
package controllers

import (
	"GoglangRevelCURD/app/models"
	"GoglangRevelCURD/app/routes"
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/viewmodels/departmentvm"

	"github.com/revel/revel"
)

type DepartmentCtrl struct {
	BaseCtrl
}

func (c DepartmentCtrl) Index() revel.Result {
	vm := departmentvm.GetAllDepartmentViewModels{
		Txn: c.Txn,
	}

	model := vm.LoadAll()
	return c.Render(model)
}
func (c DepartmentCtrl) AddDepartment() revel.Result {
	return c.Render()
}
func (c DepartmentCtrl) AddAction(department models.Department) revel.Result {
	vm := departmentvm.ActionViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		Department: department,
	}
	model := vm.AddDepartment()
	if model.IsSuccess {
		c.Flash.Success(model.Message)
	} else {
		c.Flash.Error(model.Message)
	}
	return c.Redirect(routes.DepartmentCtrl.Index())
}
func (c DepartmentCtrl) UpdateDepartment(id int) revel.Result {
	vm := departmentvm.GetByIdDepartmentViewModels{
		Txn: c.Txn,
		Id:  id,
	}

	model := vm.LoadAll()
	return c.Render(model)
}
func (c DepartmentCtrl) UpdateAction(id int, updatedepartment models.Department) revel.Result {
	vm := departmentvm.UpdateDepartmentViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		UpdateDepartment: updatedepartment,
		Id:               id,
	}
	model := vm.LoadAll()
	if model.IsSuccess {
		c.Flash.Success(model.Message)
	} else {
		c.Flash.Error(model.Message)
	}
	return c.Redirect(routes.DepartmentCtrl.Index())
}
func (c DepartmentCtrl) DeleteAction(id int) revel.Result {
	vm := departmentvm.DeleteViewModel{
		BaseVM: viewmodels.BaseVM{
			Txn: c.Txn,
		},
		Id: id,
	}

	model := vm.LoadAll()

	if model.IsSuccess == true {
		c.Flash.Success(model.Message)
	} else {
		c.Flash.Error(model.Message)
	}

	return c.Redirect(routes.DepartmentCtrl.Index())
}

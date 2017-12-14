package departmentvm

import (
	"GoglangRevelCURD/app/models"
	"GoglangRevelCURD/app/viewmodels"

	"github.com/revel/revel"
)

type UpdateDepartmentViewModel struct {
	Id int
	viewmodels.BaseVM
	UpdateDepartment models.Department
}

type UpdateDepartmentViewModelResponse struct {
	Department models.Department
	viewmodels.BaseVMResponse
}

func (vm UpdateDepartmentViewModel) LoadAll() UpdateDepartmentViewModelResponse {
	department := models.Department{}
	var errMessage string = "Gagal update Department"
	err := vm.Txn.SelectOne(&department, "SELECT * FROM `department` where id = ?", vm.Id)
	revel.WARN.Printf("%#v", err)

	department.DepartmentName = vm.UpdateDepartment.DepartmentName
	department.NumberOfEmployees = vm.UpdateDepartment.NumberOfEmployees

	success, updateErr := vm.Txn.Update(&department)

	if success > 0 && updateErr == nil {
		return UpdateDepartmentViewModelResponse{
			BaseVMResponse: viewmodels.BaseVMResponse{
				IsSuccess: true,
				Message:   "Department berhasil di Update",
			},
			Department: department,
		}
	}

	return UpdateDepartmentViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message:   errMessage,
		},
		Department: department,
	}
}

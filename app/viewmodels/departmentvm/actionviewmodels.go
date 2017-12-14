package departmentvm

import (
	"GoglangRevelCURD/app/models"
	"GoglangRevelCURD/app/viewmodels"

	"github.com/revel/revel"
)

type ActionViewModel struct {
	viewmodels.BaseVM
	Department models.Department
}

type ActionViewModelResponse struct {
	viewmodels.BaseVMResponse
	Department models.Department
}

func (vm ActionViewModel) AddDepartment() ActionViewModelResponse {
	var errMessage string = "Gagal menambahkan Department"

	err := vm.Txn.Insert(&vm.Department)
	if err == nil {
		return ActionViewModelResponse{
			BaseVMResponse: viewmodels.BaseVMResponse{
				IsSuccess: true,
				Message:   "Department berhasil di tambahkan",
			},
			Department: vm.Department,
		}
	} else {
		revel.WARN.Printf("superuser.insertUser error status %s", err.Error())
	}

	return ActionViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message:   errMessage,
		},
		Department: vm.Department,
	}
}

package departmentvm

import (
	"GoglangRevelCURD/app/models"
	"GoglangRevelCURD/app/viewmodels"

	"github.com/revel/revel"
)

type DeleteViewModel struct {
	viewmodels.BaseVM
	Id int
}

type DeleteViewModelResponse struct {
	viewmodels.BaseVMResponse
}

func (vm DeleteViewModel) LoadAll() DeleteViewModelResponse {
	department := models.Department{}
	var errMessage string = "Gagal menghapus Department"
	err := vm.Txn.SelectOne(&department, "SELECT * FROM `department` where id = ?", vm.Id)
	revel.WARN.Printf("%#v", err)
	if err == nil {
		department.Id = vm.Id
		success, updateErr := vm.Txn.Delete(&department)
		revel.WARN.Printf("%#v", updateErr)
		if success > 0 && updateErr == nil {
			return DeleteViewModelResponse{
				BaseVMResponse: viewmodels.BaseVMResponse{
					IsSuccess: true,
					Message:   "Berhasil menghapus Department",
				},
			}
		}
	}

	return DeleteViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message:   errMessage,
		},
	}
}

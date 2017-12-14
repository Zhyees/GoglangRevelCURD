package employedvm

import (
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/models"
	"github.com/revel/revel"
)

type UpdateActionViewModel struct {
	Id int
	viewmodels.BaseVM
	UpdateEmployed models.Employed
} 

type UpdateActionViewModelResponse struct {
	Employed models.Employed
	viewmodels.BaseVMResponse
}


func (vm UpdateActionViewModel) LoadAll() UpdateActionViewModelResponse{
	employed := models.Employed{}
	var errMessage string = "Gagal menambahkan Employed"
	err := vm.Txn.SelectOne(&employed, "SELECT * FROM `data` where id = ?", vm.Id)
	revel.WARN.Printf("%#v", err)

	employed.Name = vm.UpdateEmployed.Name
	employed.Email = vm.UpdateEmployed.Email
	employed.Phone = vm.UpdateEmployed.Phone
	employed.Address = vm.UpdateEmployed.Address
	employed.Photo = vm.UpdateEmployed.Photo

	success, updateErr := vm.Txn.Update(&employed)

		if success > 0 && updateErr == nil{
			return UpdateActionViewModelResponse{
			BaseVMResponse: viewmodels.BaseVMResponse{
				IsSuccess: true,
				Message: "Employed berhasil di Update",
			},
			Employed: employed,
		}
		}

	return UpdateActionViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message: errMessage,
		},
		Employed: employed,
	}
}



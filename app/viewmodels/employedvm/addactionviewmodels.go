package employedvm

import (
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/models"
	"github.com/revel/revel"
)

type AddActionViewModel struct {
	viewmodels.BaseVM
	Employed models.Employed
} 

type AddActionViewModelResponse struct {
	viewmodels.BaseVMResponse
	Employed models.Employed
}


func (vm AddActionViewModel) AddEmployed() AddActionViewModelResponse{
	var errMessage string = "Gagal menambahkan Employed"
	err := vm.Txn.Insert(&vm.Employed)
	if err == nil{
		return AddActionViewModelResponse{
			BaseVMResponse: viewmodels.BaseVMResponse{
				IsSuccess: true,
				Message: "Employed berhasil di tambahkan",
			},
			Employed: vm.Employed,
		}
	}else{
		revel.WARN.Printf("superuser.insertUser error status %s", err.Error())
	}

	return AddActionViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message: errMessage,
		},
		Employed: vm.Employed,
	}
}



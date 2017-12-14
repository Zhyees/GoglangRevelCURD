package employedvm

import (
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/models"
	"github.com/revel/revel"
)

type DeleteActionViewModel struct {
	viewmodels.BaseVM
	Id int
} 

type DeleteActionViewModelResponse struct {
	viewmodels.BaseVMResponse
}


func (vm DeleteActionViewModel) LoadAll() DeleteActionViewModelResponse{
	employed := models.Employed{}
	var errMessage string = "Gagal menghapus Employed"
	err := vm.Txn.SelectOne(&employed, "SELECT * FROM `data` where id = ?", vm.Id)
	revel.WARN.Printf("%#v", err)
	if err == nil {
		employed.Id = vm.Id
		success, updateErr := vm.Txn.Delete(&employed)
		revel.WARN.Printf("%#v", updateErr)
		if success > 0 && updateErr == nil{
			return DeleteActionViewModelResponse{
				BaseVMResponse: viewmodels.BaseVMResponse{
					IsSuccess: true,
					Message: "Berhasil menghapus Employed",
				},
			}
		}
	}
		

	return DeleteActionViewModelResponse{
		BaseVMResponse: viewmodels.BaseVMResponse{
			IsSuccess: false,
			Message: errMessage,
		},
	}
}



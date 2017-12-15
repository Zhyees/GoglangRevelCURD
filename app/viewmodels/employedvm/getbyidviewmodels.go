package employedvm

import (
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
	"GoglangRevelCURD/app/models"
)

type GetByIdViewModels struct {
	Txn *gorp.Transaction
	Id int
}

type GetByIdViewModelsResponse struct {
	Employed models.Employed
}

func (vm GetByIdViewModels) LoadAll() GetByIdViewModelsResponse{
	employed := models.Employed{}
	err := vm.Txn.SelectOne(&employed, "SELECT * FROM `employed` where id = ? ", vm.Id)
	if err == nil{
		return GetByIdViewModelsResponse{
			Employed: employed,
		}
	}

	revel.WARN.Printf("%#v", err)

	return GetByIdViewModelsResponse{}
}

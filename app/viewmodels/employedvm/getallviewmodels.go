package employedvm

import (
	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
	"GoglangRevelCURD/app/models"
)

type GetAllViewModels struct {
	Txn *gorp.Transaction
}

type GetAllViewModelsResponse struct {
	EmlpoyedList []interface{} `json:"employed"` 
}

func (vm GetAllViewModels) LoadAll() GetAllViewModelsResponse{
	employedList, err := vm.Txn.Select(models.Employed{}, "SELECT * FROM `employed`")
	if err == nil{
		return GetAllViewModelsResponse{
			EmlpoyedList: employedList,
		}
	}

	revel.WARN.Printf("%#v", err)

	return GetAllViewModelsResponse{}
}

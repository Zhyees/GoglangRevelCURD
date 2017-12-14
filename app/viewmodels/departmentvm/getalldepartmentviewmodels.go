package departmentvm

import (
	"GoglangRevelCURD/app/models"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type GetAllDepartmentViewModels struct {
	Txn *gorp.Transaction
}

type GetAllDepartmentViewModelsResponse struct {
	DepartmentList []interface{} `json:"department"`
}

func (vm GetAllDepartmentViewModels) LoadAll() GetAllDepartmentViewModelsResponse {
	departmentList, err := vm.Txn.Select(models.Department{}, "SELECT * FROM `department`")
	if err == nil {
		return GetAllDepartmentViewModelsResponse{
			DepartmentList: departmentList,
		}
	}

	revel.WARN.Printf("%#v", err)

	return GetAllDepartmentViewModelsResponse{}
}

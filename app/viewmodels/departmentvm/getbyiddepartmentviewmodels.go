package departmentvm

import (
	"GoglangRevelCURD/app/models"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
)

type GetByIdDepartmentViewModels struct {
	Txn *gorp.Transaction
	Id  int
}

type GetByIdDepartmentViewModelsResponse struct {
	Department models.Department
}

func (vm GetByIdDepartmentViewModels) LoadAll() GetByIdDepartmentViewModelsResponse {
	department := models.Department{}
	err := vm.Txn.SelectOne(&department, "SELECT * FROM `department` where id = ? ", vm.Id)
	if err == nil {
		return GetByIdDepartmentViewModelsResponse{
			Department: department,
		}
	}

	revel.WARN.Printf("%#v", err)

	return GetByIdDepartmentViewModelsResponse{}
}

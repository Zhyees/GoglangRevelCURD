package adminvm

import (
	"GoglangRevelCURD/app/models"
	"github.com/go-gorp/gorp"
)

type LoginViewModel struct {
	Superuser models.Superuser
	Txn *gorp.Transaction
}

type LoginViewModelResponse struct{
	IsLogin bool
}

func (vm LoginViewModel) LoadAll() LoginViewModelResponse{
	status := false

	superuser := vm.getSuperuserByEmail()
	if superuser.LoginToken == vm.Superuser.LoginToken{
		status = true
	}

	return LoginViewModelResponse{
		IsLogin: status,
	}
}

func (vm LoginViewModel) getSuperuserByEmail() models.Superuser{
	superuser := models.Superuser{}
	err := vm.Txn.SelectOne(&superuser, "SELECT * FROM `superuser` WHERE `email` = ?", vm.Superuser.Email)

	if err != nil{
		return models.Superuser{}
	}

	return superuser
}

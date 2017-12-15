package adminvm

import (
	"GoglangRevelCURD/app/viewmodels"
	"GoglangRevelCURD/app/models"
	"github.com/go-gorp/gorp"
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"time"
	mathRand "math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var tokenLength int = 20

type LoginActionViewModel struct {
	viewmodels.BaseVM
	Login models.Login 
	Txn *gorp.Transaction
}

type LoginActionViewModelResponse struct {
	Superuser models.Superuser `json:"superuser"`
	Role models.Role `json:"role"`
	ModulList []interface{} `json:"modullist"` 
	IsSuccess bool
}

func (vm LoginActionViewModel) LoadAll() (LoginActionViewModelResponse){
	superuser, role, err := vm.login()

	if err == nil{

		modullist, modulerr := vm.Txn.Select(models.Modul{}, "select A.* from modul AS A "+
				"left join `modul_role_mapping` AS B on A.`id` = B.`modul_id` "+
				"where B.`role_id` = ? and A.`is_deleted` = 0", superuser.RoleId)
		revel.WARN.Printf("%#v", modulerr)

		password := []byte(vm.Login.Password)
		cryptErr := bcrypt.CompareHashAndPassword([]byte(superuser.HashedPassword), password);

		if cryptErr == nil{
			response := LoginActionViewModelResponse{
				Superuser: superuser,
				Role:role,
				ModulList:modullist,
				IsSuccess:true,
			}

			return response
		}else{
			response := LoginActionViewModelResponse{
				Superuser: superuser,
				Role:role,
				ModulList:modullist,
				IsSuccess:false,
			}

			return response
		}
	}

	return LoginActionViewModelResponse{}
}


func (vm LoginActionViewModel) login() (models.Superuser, models.Role, error){
	superuser := new(models.Superuser)
	err := vm.Txn.SelectOne(superuser,
		"SELECT * FROM `superuser` WHERE `email` = ?", vm.Login.Email)

	if err == nil{

		//get role
		role := models.Role{}
		roleErr := vm.Txn.SelectOne(&role, "SELECT * FROM `role` WHERE `id` = ?", superuser.RoleId)

		if roleErr == nil{
				//generate random token
				rune := vm.randomStringRunes(tokenLength)
				superuser.LoginToken = rune

				//update login token
				success, errUpdate := vm.Txn.Update(superuser)

				if errUpdate == nil && success != 0{
					return *superuser, role, nil
				}
			


			revel.WARN.Printf("%#v", errUpdate)
		}
	}
	return models.Superuser{}, models.Role{}, err
}

func(vm LoginActionViewModel) randomStringRunes(n int) string {
	mathRand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mathRand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (vm LoginActionViewModel) updateLoginToken() {

}

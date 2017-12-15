package controllers

import (
	"GoglangRevelCURD/app/models"
	"github.com/revel/revel"
	"GoglangRevelCURD/app/routes"
	"GoglangRevelCURD/app/viewmodels/adminvm"
	"strconv"
)

/**
Admin Controller
This controller is used only for authentication transaction
No session dependent transaction is used in this controller
 */
type AdminCtrl struct {
	BaseCtrl
}

/**
Admin Login Page
 */
func (a AdminCtrl) Login() revel.Result {
	if len(a.Session["superuser"]) > 0 {
		superuser := models.Superuser{}
		err := a.fromJson(a.Session["superuser"], &superuser)
		revel.WARN.Printf("%#v", err)
		revel.WARN.Printf("%#v", superuser)
		if err != nil {
			return a.RenderError(err)
		}

		vm := adminvm.LoginViewModel{
			Superuser: superuser,
			Txn:       a.Txn,
		}

		model := vm.LoadAll()

		if model.IsLogin {
			return a.Redirect(routes.HomeCtrl.Index())
		}
	}

	return a.RenderTemplate("adminctrl/login.html")
}

/**
Admin Login Action
 */
func (a AdminCtrl) LoginAction(login models.Login) revel.Result {
	login.Validate(a.Validation)

	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		a.Flash.Error("Bad Username / Password")
		return a.Redirect(routes.AdminCtrl.Login())
	}

	vm := adminvm.LoginActionViewModel{
		Login: login,
		Txn:   a.Txn,
	}

	model := vm.LoadAll()
	if model.IsSuccess == false {
		a.Flash.Error("Bad Username / Password")
		return a.Redirect(routes.AdminCtrl.Login())
	}

	a.storeLoginSession(model.Superuser, model.Role,model.ModulList)
	return a.Redirect(routes.HomeCtrl.Index())
}

/**
Admin Logout
 */
func (a AdminCtrl) Logout() revel.Result {
	a.deleteLoginSession()
	return a.Redirect(routes.AdminCtrl.Login())
}

/**
Unauthorized Page
 */
func (a AdminCtrl) UnAuthorized() revel.Result {
	return a.RenderTemplate("adminctrl/unauthorized.html")
}

func (a AdminCtrl) storeLoginSession(superuser models.Superuser, role models.Role , modul []interface{} ) {
	superuserModelJson, err := a.toJson(superuser)
	roleModelJson, err := a.toJson(role)
	modulList, err := a.toJson(modul);
	if err == nil {
		a.Session["modul"] = modulList
		a.Session["superuser"] = superuserModelJson
		a.Session["role"] = roleModelJson
		a.Session["loginId"] = strconv.Itoa(superuser.Id)
		a.Session["loginEmail"] = superuser.Email
		a.Session["loginName"] = superuser.FullName
		a.Session["loginToken"] = superuser.LoginToken
		a.Session["loginRole"] = role.Name
		a.Session["loginRoleId"] = strconv.Itoa(role.Id)
	}

	var sidebarLayout string

	switch role.Id {
	case SuperAdminRoleId:
		sidebarLayout = "sidebar.admin.html"
	case OperationalAdminRoleId:
		sidebarLayout = "sidebar.operational.html"
	default:
		sidebarLayout = "sidebar.operational.html"
	}

	a.Session["loginSidebarLayout"] = sidebarLayout
}

func (a AdminCtrl) deleteLoginSession() {
	delete(a.Session, "superuser")
	delete(a.Session, "role")
	delete(a.Session, "loginId")
	delete(a.Session, "loginEmail")
	delete(a.Session, "loginName")
	delete(a.Session, "loginToken")
	delete(a.Session, "loginRole")
	delete(a.Session, "loginRoleId")
	delete(a.Session, "modul")
}

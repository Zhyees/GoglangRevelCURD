package controllers

import (
	"github.com/revel/revel"
	"github.com/go-gorp/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
	"GoglangRevelCURD/app/models"
	//"net/http"
	//"encoding/json"
)

func init(){

	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		//TokenAuthFilter,               // Resolve Token Authentication (Custom)
		//URLHookFilter,                 // Store current action url (Custom)
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result. [^1]
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(InitDb)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Could not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}


func getConnectionString() string {
	host := getParamString("db.host", "localhost")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "root")
	pass := getParamString("db.password", "12345abc")
	dbname := getParamString("db.name", "curd")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

var InitDb func() = func(){
	connectionString := getConnectionString()
	if db, err := sql.Open("mysql", connectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db: db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}
	// Defines the table for use by GORP
	// This is a function we will create soon.
	//defineBannerItemTable(Dbm)
	defineTables(Dbm)

	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		revel.ERROR.Fatal(err)
	}
}

func defineTables(dbm *gorp.DbMap){
	dbm.AddTableWithName(models.Employed{}, "data").SetKeys(true, "id")
}

//last effort
//do this. Create unauthorized page
//c.Result = c.RenderTemplate("admin/login.html")
//return
/*
var TokenAuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Args["controllersName"] = c.Name
	c.Session["controllersName"] = c.Name

	if c.Name != "AdminCtrl" && c.Name != "Static"{
		if len(c.Session["superuser"]) < 1{
			c.Result = c.RenderTemplate("admin/unauthorized.html")
			return
		}

		superuser := models.Superuser{}
		err := json.Unmarshal([]byte(c.Session["superuser"]), &superuser)

		if err != nil{
			c.Result = c.RenderTemplate("admin/error.html")
			return
		}

		authSuperuser := models.Superuser{}
		err = Dbm.SelectOne(&authSuperuser, "SELECT * FROM `superuser` WHERE `login_token` = ?", c.Session["loginToken"])

		//temporary disable
		//if authErr != nil{
		//	revel.WARN.Printf("%#v", c.Session["loginToken"])
		//	c.Result = c.RenderTemplate("admin/error.html")
		//	return
		//}

		var found = false

		modules, err := Dbm.Select(models.MiniApp{}, "SELECT * FROM `mini_app_map` AS A INNER JOIN " +
			"`mini_app` as B ON A.`mini_app_id` = B.`id` WHERE A.`role_id` = ?", superuser.RoleId)

		for _, moduleRef := range modules{
			module := moduleRef.(*models.MiniApp)

			if module.CtrlName == c.Name{
				found = true
			}
		}

		revel.WARN.Printf("modules length %+v", len(modules))

		if !found && len(modules) > 0 {
			c.Result = c.RenderTemplate("admin/unauthorized.html")
			return
		}
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var URLHookFilter =  func(c *revel.Controller, fc []revel.Filter) {
	if c.Name != "Static"{
		c.Session["currentUrl"] = c.Request.URL.Path
		c.ViewArgs["currentUrl"] = c.Request.URL.Path
		c.ViewArgs["currentCtrl"] = c.Name
		c.ViewArgs["currentCtrlAndMethod"] = c.Name + "/" + c.MethodName
	}

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
*/

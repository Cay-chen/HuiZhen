package routers

import (
	"HuiZhen/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/table/:id", &controllers.TableController{})
	beego.Router("/getdate", &controllers.GetDateController{})
	beego.Router("/apply", &controllers.ApplyController{})
	beego.Router("/getdate1", &controllers.GetDate1Controller{})
	//beego.Router("/table1", &controllers.Table1Controller{})
	beego.Router("/approve", &controllers.ApproveController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/register/:id", &controllers.RegisterController{})
	beego.Router("/error/:id", &controllers.ErrorController{})
	beego.Router("/exit", &controllers.ExitControoler{})
	beego.Router("/view_page", &controllers.ViewPageController{})
	beego.Router("/change_password", &controllers.ChangePasswordController{})
	beego.Router("/change_msg/:id", &controllers.ConPersonServerController{})
	beego.Router("/iframe1", &controllers.IframeController1{})
	beego.Router("/iframe2", &controllers.IframeController2{})
	beego.Router("/get_person", &controllers.GetPersonDataController{})
	beego.Router("/get_dev", &controllers.GetDevDataController{})
	beego.Router("/get_form/:id", &controllers.GetFormDataController{})
	beego.Router("/person_manage", &controllers.PersonManageController{})
	beego.Router("/dep_manage", &controllers.DepMangeController{})

}

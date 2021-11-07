package routers

import (
	"HuiZhen/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/apply", &controllers.ApplyController{})
	beego.Router("/apply_et", &controllers.ApplicationFormEtController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/register/:id", &controllers.RegisterController{})
	beego.Router("/error/:id", &controllers.ErrorController{})
	beego.Router("/exit", &controllers.ExitControoler{})
	beego.Router("/change_password", &controllers.ChangePasswordController{})
	beego.Router("/change_msg/:id", &controllers.ConPersonServerController{})
	beego.Router("/get_person", &controllers.GetPersonDataController{})
	beego.Router("/get_dev", &controllers.GetDevDataController{})
	beego.Router("/get_form/:id", &controllers.GetFormDataController{})
	beego.Router("/data/:id", &controllers.DataController{})
	beego.Router("/person_manage", &controllers.PersonManageController{})
	beego.Router("/dep_manage", &controllers.DepMangeController{})
	beego.Router("/check", &controllers.CheckController{})
	beego.Router("/approve56", &controllers.Approve5To6Controller{})
	beego.Router("/iframe/:id", &controllers.IframeController{})

}

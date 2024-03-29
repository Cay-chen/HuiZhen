package routers

import (
	"HuiZhen/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/tips", &controllers.ChromeController{})
	beego.Router("/apply", &controllers.ApplyController{})
	beego.Router("/apply_et", &controllers.ApplicationFormEtController{})
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/wj", &controllers.LoginWjController{})
	beego.Router("/register/:id", &controllers.RegisterController{})
	beego.Router("/error/:id", &controllers.ErrorController{})
	beego.Router("/exit", &controllers.ExitControoler{})
	beego.Router("/change_password", &controllers.ChangePasswordController{})
	beego.Router("/change_msg/:id", &controllers.ConPersonServerController{})
	beego.Router("/data/:id", &controllers.DataController{})
	beego.Router("/data_table/:id", &controllers.DataTableController{})
	beego.Router("/check", &controllers.CheckController{})
	beego.Router("/con_note", &controllers.ConsultationNoteController{})
	beego.Router("/approve56", &controllers.Approve5To6Controller{})
	beego.Router("/iframe/:id", &controllers.IframeController{})
}

package controllers

import (
	"HuiZhen/models/utils"
	beego "github.com/beego/beego/v2/server/web"
)

// LoginController /*登录控制器*/
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		// 显示设置成功
		c.Redirect("/index", 302)
	} else if n, ok = flash.Data["error"]; ok {
		// 显示错误
		c.Data["Err"] = "inline"
		c.Data["Err_msg"] = n
		c.TplName = "login.html"
	} else {
		// 不然默认显示设置页面
		c.Data["Err"] = "none"
		c.TplName = "login.html"
	}
}

func (c *LoginController) Post() {
	serverName := "JYConPersonServlet"
	method := "checkAccount"
	JYConPersonCode := c.GetString("JYConPersonCode")
	JYConPersonPassword := c.GetString("JYConPersonPassword")
	JYConPersonBelongHos := c.GetString("JYConPersonBelongHos")
	parameterMap := make(map[string]string)
	parameterMap["JYConPersonCode"] = JYConPersonCode
	parameterMap["JYConPersonPassword"] = utils.Md5String32(JYConPersonPassword)
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	resmsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), JYConPersonCode)
	if resmsg == "" {
		c.Redirect("error/500", 302)
		return
	}
	res := getPerson(resmsg)
	flash := beego.NewFlash()
	if res.Flag == "true" {
		err := c.SetSession("loginUser", resmsg)
		if err != nil {
			return
		}
		flash.Notice(res.Mesg)
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	} else {
		if c.GetSession("loginUser") != nil {
			err := c.DestroySession()
			if err != nil {
				return
			}
		}
		flash.Error(res.Mesg)
		flash.Store(&c.Controller)
		c.Redirect("/", 302)

	}
}

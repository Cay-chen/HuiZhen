package controllers

import (
	"HuiZhen/models/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ChangePasswordController struct {
	BaseController
}

func (c *ChangePasswordController) Get() {
	if c.IsLogin {
		c.Data["JYConPersonName"] = c.PersonUer.JYConPersonName
		c.Data["JYConPersonCode"] = c.PersonUer.JYConPersonCode
		flash := beego.ReadFromRequest(&c.Controller)
		if n, ok := flash.Data["notice"]; ok {
			fmt.Println("notice:" + n)
			c.Data["Msg_ture"] = true
			// 显示设置成功
			c.Data["Msg_notice"] = n
		} else if n, ok = flash.Data["error"]; ok {
			c.Data["Msg_ture"] = true
			// 显示错误
			c.Data["Msg_notice"] = n
		} else {
			// 不然默认显示设置页面
			c.Data["Msg_notice"] = ""
		}
		c.TplName = "change_password.html"
	} else {
		c.TplName = "error.html"
	}

}
func (c *ChangePasswordController) Post() {
	if c.IsLogin {
		JYConPersonNewPassword := c.GetString("JYConPersonNewPassword")
		JYConPersonOldPassword := c.GetString("JYConPersonOldPassword")
		serverName := "JYConPersonServlet"
		method := "changeAccountPassword"
		parameterMap := make(map[string]string)
		parameterMap["JYConPersonNewPassword"] = utils.Md5String32(JYConPersonNewPassword)
		parameterMap["JYConPersonOldPassword"] = utils.Md5String32(JYConPersonOldPassword)
		parameterMap["JYConPersonCode"] = c.PersonUer.JYConPersonCode
		parameterMap["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
		resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
		res := getPerson(resMsg)
		flash := beego.NewFlash()
		if res.Flag == "true" {
			flash.Notice(res.Mesg)
		} else {
			flash.Error(res.Mesg)
		}
		flash.Store(&c.Controller)
		c.Redirect("/change_password", 302)
	} else {
		c.TplName = "error.html"
	}

}

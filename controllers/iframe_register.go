package controllers

import (
	"HuiZhen/models"
	"HuiZhen/models/utils"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	BaseController
}

func (c *RegisterController) Get() {
	if c.IsLogin {
		if !c.IsYxys {
			id := c.Ctx.Input.Param(":id")
			switch id {
			case "person":
				if c.IsKszr || c.IsFzr {
					res := "[{\"JYConDepCode\": \"" + c.PersonUer.JYConPersonBelongDep + "\", \"JYConDepName\": \"" + c.PersonUer.DepName + "\"}]"
					var res1 []models.Dep
					_ = json.Unmarshal([]byte(res), &res1)
					c.Data["DepList"] = res1
					c.Data["IsYwb"] = false

				} else {
					c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200", c.PersonUer.JYConPersonCode).Data
					c.Data["IsYwb"] = true
				}
				flash := beego.ReadFromRequest(&c.Controller)
				if n, ok := flash.Data["notice"]; ok {
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
				c.TplName = "iframe_register_person.html"
				break
			case "dep":
				flash := beego.ReadFromRequest(&c.Controller)
				if n, ok := flash.Data["notice"]; ok {
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
				c.TplName = "iframe_register_dep.html"
				break
			}
		} else {
			c.Redirect("/error", 302)
		}

	} else {
		c.Redirect("/error", 302)
	}

}
func (c *RegisterController) Post() {
	if c.IsLogin {
		if !c.IsYxys {
			id := c.Ctx.Input.Param(":id")
			switch id {
			case "person":
				JYConPersonCode := c.GetString("JYConPersonCode")
				JYConPersonPassword := c.GetString("JYConPersonPassword")
				JYConPersonName := c.GetString("JYConPersonName")
				JYConPersonType := c.GetString("JYConPersonType")
				JYConPersonBelongDep := c.GetString("JYConPersonBelongDep")
				JYConPersonPhone := c.GetString("JYConPersonPhone")
				serverName := "JYConPersonServlet"
				method := "createAccount"
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonCode"] = JYConPersonCode
				parameterMap["JYConPersonPassword"] = utils.Md5String32(JYConPersonPassword)
				parameterMap["JYConPersonName"] = JYConPersonName
				parameterMap["JYConPersonType"] = JYConPersonType
				parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
				parameterMap["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
				parameterMap["JYConPersonPhone"] = JYConPersonPhone
				resmsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				res := getMsg(resmsg)
				flash := beego.NewFlash()
				if res.Flag == "true" {
					flash.Notice(res.Mesg)
				} else {
					flash.Error(res.Mesg)
				}
				flash.Store(&c.Controller)
				c.Redirect("/register/person", 302)
				break
			case "dep":
				JYConDepCode := c.GetString("JYConDepCode")
				JYConDepName := c.GetString("JYConDepName")
				JYConDepPhone := c.GetString("JYConDepPhone")
				JYConDepLocalhost := c.GetString("JYConDepLocalhost")
				JYConDepIsActive := "0"
				if c.GetString("JYConDepIsActive") == "on" {
					JYConDepIsActive = "1"
				}
				serverName := "JYConDepListServlet"
				method := "createDep"
				parameterMap := make(map[string]string)
				parameterMap["JYConDepCode"] = JYConDepCode
				parameterMap["JYConDepName"] = JYConDepName
				parameterMap["JYConDepPhone"] = JYConDepPhone
				parameterMap["JYConDepLocalhost"] = JYConDepLocalhost
				parameterMap["JYConDepBelongHos"] = c.PersonUer.JYConPersonBelongHos
				parameterMap["JYConDepIsActive"] = JYConDepIsActive
				resmsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				res := getMsg(resmsg)
				flash := beego.NewFlash()
				if res.Flag == "true" {
					flash.Notice(res.Mesg)
				} else {
					flash.Error(res.Mesg)
				}
				flash.Store(&c.Controller)
				c.Redirect("/register/dep", 302)
				break
			}
		} else {
			c.Redirect("/error", 302)
		}

	} else {
		c.Redirect("/error", 302)
	}

}

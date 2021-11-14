package controllers

import (
	"HuiZhen/models"
	"HuiZhen/models/utils"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type ConsultationNoteController struct {
	BaseController
}

func (c *ConsultationNoteController) Get() {
	if c.IsLogin {
		flash := beego.ReadFromRequest(&c.Controller)
		if n, ok := flash.Data["notice"]; ok {
			// 显示设置成功
			c.Data["SpOk"] = true
			c.TplName = "iframe_consultation_note.html"
		} else if n, ok = flash.Data["error"]; ok {
			// 显示错误
			c.Data["SpOk"] = true
			c.Data["Err_msg"] = n
			c.TplName = "iframe_consultation_note.html"
		} else {
			c.Data["SpOk"] = false
			JYConNum := c.GetString("JYConNum")
			if JYConNum != "" {
				info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "", c.PersonUer.JYConPersonCode)
				s := models.FormInfo{}
				err := json.Unmarshal([]byte(info), &s)
				if err != nil {
				}
				c.Data["Form"] = s
				c.TplName = "iframe_consultation_note.html"
			} else {
				c.Redirect("/error/600", 302)
			}
		}

	} else {
		c.Redirect("/error/600", 302)
	}

}
func (c *ConsultationNoteController) Post() {
	if c.IsLogin {
		JYConNum := c.GetString("JYConNum")
		JYConFormPolicy := c.GetString("JYConFormPolicy")
		JYConFormConclusion := c.GetString("JYConFormConclusion")
		JYConFormConclusionDate := c.GetString("JYConFormConclusionDate")
		JYConFormApproveComment := "写会诊记录"
		serverName := "JYConFormServlet"
		method := "changeFormState"
		parameterMap := make(map[string]string)
		parameterMap["JYConNum"] = JYConNum
		parameterMap["JYConFormConclusion"] = JYConFormConclusion
		parameterMap["JYConFormConclusionDate"] = JYConFormConclusionDate
		parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
		parameterMap["flag"] = "approve"
		parameterMap["JYConNum"] = JYConNum
		parameterMap["JYConType"] = "1"
		parameterMap["JYConFormPolicy"] = JYConFormPolicy
		parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
		parameterMap["JYConFormApprovePersonId"] = c.PersonUer.JYConPersonCode
		parameterMap["JYConFormApprovePersonName"] = c.PersonUer.JYConPersonName
		parameterMap["JYConFormApproveBelongHos"] = c.PersonUer.JYConPersonBelongHos
		resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
		result := getPerson(resMsg)
		flash := beego.NewFlash()
		/*	if result.Flag == "true" {
				flash.Notice(result.Mesg)
				flash.Store(&c.Controller)
				c.Redirect("/index", 302)
			} else {

			}*/

		if result.Flag == "true" {
			flash.Notice(result.Mesg)
			flash.Store(&c.Controller)
			c.Redirect("/con_note", 302)
		} else {
			flash.Error(result.Mesg)
			flash.Store(&c.Controller)
			c.Redirect("/con_note", 302)

		}

	}

}

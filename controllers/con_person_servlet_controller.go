package controllers

import (
	"HuiZhen/models/utils"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
)

type ConPersonServerController struct {
	BaseController
}

func (c *ConPersonServerController) Get() {

	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		logs.Debug("id:" + id)
		switch id {
		case "deleteForm":
			serverName := "JYConFormServlet"
			method := "deleteForm"
			JYConNum := c.GetString("JYConNum")
			parameterMap := make(map[string]string)
			parameterMap["JYConNum"] = JYConNum
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			result := getPerson(resMsg)
			if result.Flag == "true" {
				c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				return
			} else {
				c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				return
			}
			break
		case "changeFormState":
			serverName := "JYConFormServlet"
			method := "changeFormState"
			flag := c.GetString("flag")
			JYConOppDocId := c.GetString("JYConOppDocId")               //设定的会诊医生ID--在状态为4提升时需要
			JYConOppDocName := c.GetString("JYConOppDocName")           // 设定的会诊医生姓名--在状态为4提升时需要
			JYConOppDocPhone := c.GetString("JYConOppDocPhone")         //设定的会诊医生电话--在状态为4提升时需要
			JYConPersonBelongDep := c.GetString("JYConPersonBelongDep") //拟邀科室代号--在状态为4提升时需要
			JYConPersonBelongHos := c.GetString("JYConPersonBelongHos") //拟邀科室院区--在状态为4提升时需要
			JYConFormApproveComment := c.GetString("JYConFormApproveComment")
			JYConFormApprovePersonId := c.PersonUer.JYConPersonCode
			JYConFormApprovePersonName := c.PersonUer.JYConPersonName
			JYConFormApproveBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConNum := c.GetString("JYConNum")
			parameterMap := make(map[string]string)
			parameterMap["JYConNum"] = JYConNum
			parameterMap["flag"] = flag
			parameterMap["JYConOppDocId"] = JYConOppDocId
			parameterMap["JYConOppDocName"] = JYConOppDocName
			parameterMap["JYConOppDocPhone"] = JYConOppDocPhone
			parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
			parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
			parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
			parameterMap["JYConFormApprovePersonId"] = JYConFormApprovePersonId
			parameterMap["JYConFormApprovePersonName"] = JYConFormApprovePersonName
			parameterMap["JYConFormApproveBelongHos"] = JYConFormApproveBelongHos
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			result := getPerson(resMsg)
			if result.Flag == "true" {
				c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				return
			} else {
				c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				return
			}
			break
		default:
			c.TplName = "500.html"
			break

		}

		if c.IsAdmin || c.IsYwb {
			switch id {
			case "active_person":
				serverName := "JYConPersonServlet"
				method := "changeAccountActive"
				JYConPersonCode := c.GetString("JYConPersonCode")
				action := c.GetString("action")
				JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonCode"] = JYConPersonCode
				parameterMap["action"] = action
				parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				fmt.Println(result)
				break
			case "cp":
				serverName := "JYConPersonServlet"
				method := "defaultPersonPassword"
				PersonNewPassword := utils.Md5String32("0")
				JYConPersonCode := c.GetString("JYConPersonCode")
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonCode"] = JYConPersonCode
				parameterMap["PersonNewPassword"] = PersonNewPassword
				parameterMap["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				fmt.Println(result)
				break
			case "change_person":
				serverName := "JYConPersonServlet"
				method := "changeAccountInfo"
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonCode"] = c.GetString("JYConPersonCode")
				parameterMap["JYConPersonName"] = c.GetString("JYConPersonName")
				parameterMap["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
				parameterMap["JYConPersonPhone"] = c.GetString("JYConPersonPhone")
				parameterMap["JYConPersonType"] = c.GetString("JYConPersonType")
				parameterMap["JYConPersonBelongDep"] = c.GetString("JYConPersonBelongDep")
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				break
			case "change_dep":
				serverName := "JYConDepListServlet"
				method := "changeDepInfo"
				parameterMap := make(map[string]string)
				parameterMap["JYConDepCode"] = c.GetString("JYConDepCode")
				parameterMap["JYConDepName"] = c.GetString("JYConDepName")
				parameterMap["JYConDepBelongHos"] = c.PersonUer.JYConPersonBelongHos
				parameterMap["JYConDepPhone"] = c.GetString("JYConDepPhone")
				parameterMap["JYConDepLocalhost"] = c.GetString("JYConDepLocalhost")
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				break
			case "active_dep":
				serverName := "JYConDepListServlet"
				method := "changeDepActive"
				JYConDepCode := c.GetString("JYConDepCode")
				action := c.GetString("action")
				JYConDepBelongHos := c.PersonUer.JYConPersonBelongHos
				parameterMap := make(map[string]string)
				parameterMap["JYConDepCode"] = JYConDepCode
				parameterMap["action"] = action
				parameterMap["JYConDepBelongHos"] = JYConDepBelongHos
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				fmt.Println(result)
				break
			default:
				c.TplName = "500.html"
				break
			}
		}
	} else {
		//返回错误
	}

}

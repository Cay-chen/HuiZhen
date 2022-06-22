package controllers

import (
	"HuiZhen/models/utils"
	"github.com/beego/beego/v2/adapter/logs"
)

type ConPersonServerController struct {
	BaseController
}

func (c *ConPersonServerController) Get() {

	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "deleteForm":
			serverName := "JYConFormServlet"
			method := "deleteForm"
			JYConNum := c.GetString("JYConNum")
			parameterMap := make(map[string]string)
			parameterMap["JYConNum"] = JYConNum
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
			result := getPerson(resMsg)
			if result.Flag == "true" {
				c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")

			} else {
				c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")

			}
			break
		case "zuofei":
			logs.Debug("")
			serverName := "JYConFormServlet"
			method := "invalidForm"
			JYConNum := c.GetString("JYConNum")
			JYConFormApproveComment := c.GetString("JYConFormApproveComment")
			JYConFormPolicy := c.GetString("JYConFormPolicy")
			JYConFormApprovePersonId := c.PersonUer.JYConPersonCode
			JYConFormApprovePersonName := c.PersonUer.JYConPersonName
			JYConFormApproveBelongHos := c.PersonUer.JYConPersonBelongHos
			parameterMap := make(map[string]string)
			parameterMap["JYConNum"] = JYConNum
			parameterMap["JYConFormPolicy"] = JYConFormPolicy
			parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
			parameterMap["JYConFormApprovePersonId"] = JYConFormApprovePersonId
			parameterMap["JYConFormApprovePersonName"] = JYConFormApprovePersonName
			parameterMap["JYConFormApproveBelongHos"] = JYConFormApproveBelongHos
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
			result := getPerson(resMsg)
			if result.Flag == "true" {
				c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")

			} else {
				c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")

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
			JYConFormConclusion := c.GetString("JYConFormConclusion")
			JYConFormPolicy := c.GetString("JYConFormPolicy")
			JYConType := c.GetString("JYConType")
			JYConFormApprovePersonId := c.PersonUer.JYConPersonCode
			JYConFormApprovePersonName := c.PersonUer.JYConPersonName
			JYConFormApproveBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConNum := c.GetString("JYConNum")
			parameterMap := make(map[string]string)
			parameterMap["JYConNum"] = JYConNum
			parameterMap["flag"] = flag
			parameterMap["JYConOppDocId"] = JYConOppDocId
			parameterMap["JYConType"] = JYConType
			parameterMap["JYConFormPolicy"] = JYConFormPolicy
			parameterMap["JYConOppDocName"] = JYConOppDocName
			parameterMap["JYConOppDocPhone"] = JYConOppDocPhone
			parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
			parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
			parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
			parameterMap["JYConFormApprovePersonId"] = JYConFormApprovePersonId
			parameterMap["JYConFormApprovePersonName"] = JYConFormApprovePersonName
			parameterMap["JYConFormApproveBelongHos"] = JYConFormApproveBelongHos
			parameterMap["JYConFormConclusion"] = JYConFormConclusion
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
			result := getPerson(resMsg)
			if result.Flag == "true" {
				c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")

			} else {
				c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")

			}
			break
		default:
			c.TplName = "600.html"
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
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				break
				/*			case "zuofei":
							logs.Debug("")
							serverName := "JYConFormServlet"
							method := "invalidForm"
							JYConNum := c.GetString("JYConNum")
							JYConFormApproveComment := c.GetString("JYConFormApproveComment")
							JYConFormPolicy := c.GetString("JYConFormPolicy")
							logs.Debug("JYConFormApproveComment")
							JYConFormApprovePersonId := c.PersonUer.JYConPersonCode
							JYConFormApprovePersonName := c.PersonUer.JYConPersonName
							JYConFormApproveBelongHos := c.PersonUer.JYConPersonBelongHos
							parameterMap := make(map[string]string)
							parameterMap["JYConNum"] = JYConNum
							parameterMap["JYConFormPolicy"] = JYConFormPolicy
							parameterMap["JYConFormApproveComment"] = JYConFormApproveComment
							parameterMap["JYConFormApprovePersonId"] = JYConFormApprovePersonId
							parameterMap["JYConFormApprovePersonName"] = JYConFormApprovePersonName
							parameterMap["JYConFormApproveBelongHos"] = JYConFormApproveBelongHos
							resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
							result := getPerson(resMsg)
							if result.Flag == "true" {
								c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")

							} else {
								c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")

							}
							break*/
			case "cp":
				serverName := "JYConPersonServlet"
				method := "defaultPersonPassword"
				PersonNewPassword := utils.Md5String32("0")
				JYConPersonCode := c.GetString("JYConPersonCode")
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonCode"] = JYConPersonCode
				parameterMap["PersonNewPassword"] = PersonNewPassword
				parameterMap["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
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
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
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
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
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
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				result := getPerson(resMsg)
				if result.Flag == "true" {
					c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
				} else {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
				}
				break
			case "add_extra":
				serverName := "JYConPersonExtraInfoServlet"
				method := "setExtraPersonInfo"
				JYConPersonAutoNum := c.GetString("JYConPersonAutoNum")
				JYConPersonCode := c.GetString("JYConPersonCode")
				JYConPersonName := c.GetString("JYConPersonName")
				JYConPersonType := c.GetString("JYConPersonType")
				JYConPersonBelongDep := c.GetString("JYConPersonBelongDep")
				JYConPersonPhone := c.GetString("JYConPersonPhone")
				JYConPersonFromNum := c.GetString("JYConPersonFromNum")
				JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
				parameterMap := make(map[string]string)
				parameterMap["JYConPersonAutoNum"] = JYConPersonAutoNum
				parameterMap["JYConPersonCode"] = JYConPersonCode
				parameterMap["JYConPersonName"] = JYConPersonName
				parameterMap["JYConPersonType"] = JYConPersonType
				parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
				parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
				parameterMap["JYConPersonPhone"] = JYConPersonPhone
				parameterMap["JYConPersonFromNum"] = JYConPersonFromNum
				resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
				result := getPerson(resMsg)
				if resMsg == "" {
					c.Ctx.WriteString("{\"code\": 1,\"msg\": \"服务器异常！\"}")
				} else {
					if result.Flag == "true" {
						c.Ctx.WriteString("{\"code\": 0,\"msg\": \"" + result.Mesg + "\"}")
					} else {
						c.Ctx.WriteString("{\"code\": 1,\"msg\": \"" + result.Mesg + "\"}")
					}
				}

				break
			default:
				c.TplName = "600.html"
				break
			}
		}
	} else {
		//返回错误
	}

}

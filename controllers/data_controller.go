package controllers

import (
	"HuiZhen/models"
	"HuiZhen/models/utils"
	"encoding/json"
)

type DataController struct {
	BaseController
}

func (c *DataController) Get() {
	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "depLocation":
			JYConDepCode := c.GetString("JYConDepCode")
			serverName := "JYConDepListServlet"
			method := "queryDepInfo"
			parameterMap := make(map[string]string)
			parameterMap["JYConDepCode"] = JYConDepCode
			parameterMap["JYConDepBelongHos"] = c.PersonUer.JYConPersonBelongHos
			res := utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
			c.Ctx.WriteString(res)
			break
		case "depList":
			JYConDepBelongHos := c.GetString("JYConDepBelongHos")
			data := models.GetDepList(JYConDepBelongHos, "1", "200", c.PersonUer.JYConPersonCode)
			res, _ := json.Marshal(data)
			c.Ctx.WriteString(string(res))
			break
		case "docList":
			JYConDepCode := c.GetString("JYConDepCode")
			yq := c.GetString("yq")
			JYConPersonBelongHos := ""
			if yq != "" {
				JYConPersonBelongHos = yq
			} else {
				JYConPersonBelongHos = c.PersonUer.JYConPersonBelongHos
			}
			res := models.GetDocList(JYConPersonBelongHos, JYConDepCode, "", "200", "1", c.PersonUer.JYConPersonCode)
			c.Ctx.WriteString(res)
			break
		case "docInfo":
			JYConPersonCode := c.GetString("JYConPersonCode")
			yq := c.GetString("yq")
			JYConPersonBelongHos := ""
			if yq != "" {
				JYConPersonBelongHos = yq
			} else {
				JYConPersonBelongHos = c.PersonUer.JYConPersonBelongHos
			}
			res := models.GetDocInfo(JYConPersonCode, JYConPersonBelongHos, c.PersonUer.JYConPersonCode)
			//res := models.GetDocList(c.PersonUer.JYConPersonBelongHos, JYConDepCode,"",  "200", "1")
			c.Ctx.WriteString(res)
			break
		default:
			c.Ctx.WriteString("{\"JYConDepLocalhost\":\"\",\"flag\":\"false\",\"JYConDepCode\":\"\",\"JYConDepName\":\"\",\"JYConDepBelongHos\":\"\",\"JYConDepPhone\":\"\",\"mesg\":\"没有数据\"}")
			break
		}
	}
}

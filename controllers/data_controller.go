package controllers

import (
	"HuiZhen/models/date"
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
			res := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			c.Ctx.WriteString(res)
			break
		case "depList":
			JYConDepBelongHos := c.GetString("JYConDepBelongHos")
			data := date.GetDepList(JYConDepBelongHos, "1", "200")
			res, _ := json.Marshal(data)
			c.Ctx.WriteString(string(res))
		default:
			c.Ctx.WriteString("{\"JYConDepLocalhost\":\"\",\"flag\":\"false\",\"JYConDepCode\":\"\",\"JYConDepName\":\"\",\"JYConDepBelongHos\":\"\",\"JYConDepPhone\":\"\",\"mesg\":\"没有数据\"}")

		}
	}
}

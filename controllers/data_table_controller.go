package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type DataTableController struct {
	BaseController
}

func (c *DataTableController) Get() {

	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "to_my":
			method := "getFormBySomethingTo"
			page := c.GetString("page")
			limit := c.GetString("limit")
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			res := models.GetFormList("", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", method, "")
			c.Ctx.WriteString(res)
			break
		case "extra":
			page := c.GetString("page")
			limit := c.GetString("limit")
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			res := models.GetExtraInfoList(JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit)
			body, _ := json.Marshal(res)
			c.Ctx.WriteString(string(body))
			break
		}

	} else {
		c.Redirect("/error/405", 302)
	}

}

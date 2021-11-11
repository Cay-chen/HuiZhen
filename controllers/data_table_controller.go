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
		case "approve":
			JYConNum := c.GetString("JYConNum")
			page := c.GetString("page")
			limit := c.GetString("limit")
			res := models.GetApproveList(JYConNum, limit, page)
			body, _ := json.Marshal(res)
			c.Ctx.WriteString(string(body))
			break
		case "person":
			if !c.IsYxys {
				limit := c.GetString("limit")
				page := c.GetString("page")
				JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
				JYConDepCode := c.GetString("key[JYConDepCode]")
				keshi := c.GetString("keshi")
				JYConPersonType := ""
				if keshi == "ks" {
					JYConPersonType = c.PersonUer.JYConPersonType
					JYConDepCode = c.PersonUer.JYConPersonBelongDep
				}
				res := models.GetDocList(JYConPersonBelongHos, JYConDepCode, JYConPersonType, limit, page)
				if res != "" {
					c.Ctx.WriteString(res)
				} else {
					c.Redirect("/error/600", 302)
				}
			}

			break

		}

	} else {
		c.Redirect("/error/405", 302)
	}

}

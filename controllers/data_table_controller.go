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
			res := models.GetFormList("", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", method, "", c.PersonUer.JYConPersonCode)
			c.Ctx.WriteString(res)
			break
		case "extra":
			page := c.GetString("page")
			limit := c.GetString("limit")
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			res := models.GetExtraInfoList(JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, c.PersonUer.JYConPersonCode)
			body, _ := json.Marshal(res)
			c.Ctx.WriteString(string(body))
			break
		case "approve":
			JYConNum := c.GetString("JYConNum")
			page := c.GetString("page")
			limit := c.GetString("limit")
			res := models.GetApproveList(JYConNum, limit, page, c.PersonUer.JYConPersonCode)
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
				res := models.GetDocList(JYConPersonBelongHos, JYConDepCode, JYConPersonType, limit, page, c.PersonUer.JYConPersonCode)
				if res != "" {
					c.Ctx.WriteString(res)
				} else {
					c.Redirect("/error/600", 302)
				}
			}
			break
		case "my":
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			res := models.GetFormList("many", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", "getFormBySomething", "", c.PersonUer.JYConPersonCode)
			c.Ctx.WriteString(res)
			break
		case "fromDsp":
			limit := c.GetString("limit")
			page := c.GetString("page")
			type1 := c.GetString("type")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			res := models.GetFormList("", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", "getMyApprovalForm", type1, c.PersonUer.JYConPersonCode)
			c.Ctx.WriteString(res)
			break
		case "fromYsp":
			limit := c.GetString("limit")
			page := c.GetString("page")
			type1 := c.GetString("type")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConPersonCode := c.PersonUer.JYConPersonCode
			res := models.GetPersonApproveList(JYConPersonCode, JYConPersonBelongHos, type1, limit, page, c.PersonUer.JYConPersonCode)
			body, _ := json.Marshal(res)
			backRes := string(body)
			c.Ctx.WriteString(backRes)
			break
		case "dep":
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			models.GetDepList(JYConPersonBelongHos, page, limit, c.PersonUer.JYConPersonCode)
			body, _ := json.Marshal(models.GetDepList(JYConPersonBelongHos, page, limit, c.PersonUer.JYConPersonCode))
			c.Ctx.WriteString(string(body))
			break
		}

	} else {
		c.Redirect("/error/405", 302)
	}

}

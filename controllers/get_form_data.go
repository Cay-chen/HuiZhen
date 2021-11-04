package controllers

import (
	"HuiZhen/models"
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
)

//设置返回类型为string

type GetFormDataController struct {
	BaseController
}

func (c *GetFormDataController) Get() {
	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "my":
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConPersonType := c.PersonUer.JYConPersonType
			JYConPersonBelongDep := c.PersonUer.JYConPersonBelongDep
			JYConPersonCode := c.PersonUer.JYConPersonCode
			res := models.GetFormList("many", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", "getFormBySomething", "")
			logs.Debug("form：" + res)
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
			res := models.GetFormList("", JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, "", "getMyApprovalForm", type1)
			c.Ctx.WriteString(res)
			break
		case "fromYsp":
			limit := c.GetString("limit")
			page := c.GetString("page")
			type1 := c.GetString("type")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConPersonCode := c.PersonUer.JYConPersonCode
			res := models.GetPersonApproveList(JYConPersonCode, JYConPersonBelongHos, type1, limit, page)
			body, _ := json.Marshal(res)
			backRes := string(body)
			c.Ctx.WriteString(backRes)
			break
		}

	} else {
		c.Redirect("/error/5434", 302)
	}

}

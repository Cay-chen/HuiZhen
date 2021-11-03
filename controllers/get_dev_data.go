package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

//设置返回类型为string

type GetDevDataController struct {
	BaseController
}

func (c *GetDevDataController) Get() {
	if c.IsLogin {
		if c.IsAdmin || c.IsYwb {
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			models.GetDepList(JYConPersonBelongHos, page, limit)
			body, _ := json.Marshal(models.GetDepList(JYConPersonBelongHos, page, limit))
			c.Ctx.WriteString(string(body))
		}
	} else {
		c.Redirect("/error/5434", 302)
	}

}

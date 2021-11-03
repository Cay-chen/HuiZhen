package controllers

import (
	"HuiZhen/models"
)

type GetPersonDataController struct {
	BaseController
}

func (c *GetPersonDataController) Get() {
	if c.IsLogin {
		if c.IsAdmin || c.IsYwb {
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConDepCode := c.GetString("key[JYConDepCode]")
			JYConPersonCode := c.GetString("key[JYConPersonCode]")
			res := models.GetDocList(JYConPersonBelongHos, JYConDepCode, JYConPersonCode, limit, page)
			if res != "" {
				c.Ctx.WriteString(res)
			} else {
				c.Redirect("/error/5434", 302)
			}
		} else {
			c.Redirect("/error/5434", 302)
		}

	} else {
		c.Redirect("/error/5434", 302)
	}

}

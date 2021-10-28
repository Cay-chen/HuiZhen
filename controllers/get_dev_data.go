package controllers

import (
	"HuiZhen/models/date"
	"encoding/json"
	"fmt"
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
			date.GetDepList(JYConPersonBelongHos, page, limit)
			body, _ := json.Marshal(date.GetDepList(JYConPersonBelongHos, page, limit))
			fmt.Println("ABCDE" + string(body))

			c.Ctx.WriteString(string(body))
		}
	} else {
		c.Redirect("/error/5434", 302)
	}

}

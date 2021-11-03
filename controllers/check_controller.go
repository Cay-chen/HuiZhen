package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type CheckController struct {
	BaseController
}

func (c *CheckController) Get() {
	if c.IsLogin {

		JYConNum := c.GetString("JYConNum")
		if JYConNum != "" {
			info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "")
			s := models.FormInfo{}
			err := json.Unmarshal([]byte(info), &s)
			if err != nil {
			}
			c.Data["Form"] = s
			c.TplName = "view_page.html"
		} else {
			c.Redirect("/error", 302)
		}

	} else {
		c.Ctx.WriteString("")
	}

}

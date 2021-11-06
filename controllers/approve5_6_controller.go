package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type Approve5To6Controller struct {
	BaseController
}

func (c *Approve5To6Controller) Get() {
	if c.IsLogin {
		JYConNum := c.GetString("JYConNum")
		c.Data["JYConNum"] = JYConNum
		c.Data["JYConPersonBelongDep"] = c.PersonUer.JYConPersonBelongDep
		c.Data["JYConPersonBelongHos"] = c.PersonUer.JYConPersonBelongHos
		docList := models.GetDocList(c.PersonUer.JYConPersonBelongHos, c.PersonUer.JYConPersonBelongDep, "", "200", "1")
		res := models.Person{}
		_ = json.Unmarshal([]byte(docList), &res)
		c.Data["DocList"] = res.Data
		c.TplName = "approve5_6.html"
	} else {
		c.Redirect("/error", 302)
	}

	//JYConDepCode（科室代码）
	//	JYConDepBelongHos

}

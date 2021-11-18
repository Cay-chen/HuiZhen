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
		docList := models.GetDocList2(c.PersonUer.JYConPersonBelongDep, c.PersonUer.JYConPersonBelongHos, "200", "1", c.PersonUer.JYConPersonCode)
		res := models.Person{}
		_ = json.Unmarshal([]byte(docList), &res)
		c.Data["DocList"] = res.Data
		c.TplName = "iframe_approve5to6.html"
	} else {
		c.Redirect("/error", 302)
	}

	//JYConDepCode（科室代码）
	//	JYConDepBelongHos

}

package controllers

import "C"
import "HuiZhen/models"

type PersonManageController struct {
	BaseController
}

func (c *PersonManageController) Get() {

	if c.IsLogin {
		if c.IsAdmin || c.IsYwb {
			c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200").Data
			c.Data["A11"] = "{{#"
			c.Data["A12"] = "{ }}"
			c.Data["A2"] = "{{d.JYConPersonIsActive}}"
			c.Data["A3"] = "{{#  } else { }}\n    {{d.JYConPersonIsActive}}\n    {{#  } }}"
			c.TplName = "iframe_person_manage.html"
		}
	} else {
		c.Redirect("/error/1245", 302)
	}

}

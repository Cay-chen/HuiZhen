package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	if c.IsLogin {
		c.Data["HuiZhen"] = "我的会诊"
		if c.IsFzr || c.IsKszr || c.IsYwb || c.IsAdmin {
			c.Data["HuiZhen"] = "科室会诊"
			c.Data["IsFzrAndZr"] = true
		} else {
			c.Data["IsFzrAndZr"] = false
		}

		if c.IsYwb || c.IsAdmin {
			c.Data["HuiZhen"] = "全院会诊"
			c.Data["IsYwbAndAdmin"] = true
		} else {
			c.Data["IsYwbAndAdmin"] = false
		}
		c.Data["PersonName"] = c.PersonUer.JYConPersonName
		switch c.PersonUer.JYConPersonBelongHos {
		case "WJQY":
			c.Data["hospital"] = "成都市温江区人民医院"
			break
		case "JYZX":
			c.Data["hospital"] = "四川省精神医学中心"
			break
		}
		c.TplName = "index.html"
	} else {
		c.Redirect("/", 302)
	}
}

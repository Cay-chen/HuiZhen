package controllers

import (
	"fmt"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	fmt.Println(c.IsLogin)
	if c.IsLogin {
		/*List1 :=false
		List2 :=false
		List3 :=false
		List4 :=false
		List0 :=false
		switch c.PersonUer.JYConPersonType {
		case "0":
			//jyConPersonType = "管理员"
			List0 =true
			break
		case "1":
			//jyConPersonType = "科室主任"
			List1 =true

			break
		case "2":
			//jyConPersonType = "科室负责人"
			List2 =true
			break
		case "3":
			//jyConPersonType = "科室一线医生"
			List3 =true
			break
		case "4":
			//jyConPersonType = "医务部人员"
			List4 =true
			break
		}*/

		c.Data["List4"] = c.IsYwb
		c.Data["List3"] = c.IsYxys
		c.Data["List2"] = c.IsFzr
		c.Data["List1"] = c.IsKszr
		c.Data["List0"] = c.IsAdmin
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
		c.Ctx.Redirect(302, "/")
	}
}

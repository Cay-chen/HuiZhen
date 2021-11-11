package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Get() {
	id := c.Ctx.Input.Param(":id")
	switch id {
	case "600":
		c.TplName = "600.html"
		break
	case "403":
		c.TplName = "403.html"
		break
	case "400":
		c.TplName = "400.html"
		break
	case "405":
		c.TplName = "405.html"
		break
	case "407":
		c.TplName = "407.svg"
		break
	case "500":
		c.TplName = "500.svg"
		break
	default:
		c.TplName = "500.html"
		break

	}

}

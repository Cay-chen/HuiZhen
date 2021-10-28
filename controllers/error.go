package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Get() {
	id := c.Ctx.Input.Param(":id")
	fmt.Println("参数：" + id)
	switch id {
	case "400":
		c.TplName = "400.html"
		break
	case "error":
		c.TplName = "error.html"
		break
	case "401":
		c.TplName = "401.html"
		break
	case "402":
		c.TplName = "402.html"
		break
	case "svg":
		c.TplName = "error.svg"
		break
	default:
		c.TplName = "500.html"
		break

	}
	//c.TplName="error.html"

}

/*func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.html"
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.html"
}


func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "600.html"
}*/

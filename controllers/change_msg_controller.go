package controllers

import "fmt"

type ChangeMsgController struct {
	BaseController
}

func (c *ChangeMsgController) Get() {

	abc := c.GetString("abc")
	def := c.GetString("def")
	fmt.Println(abc + ":" + def)
	c.Ctx.WriteString("{ \" abc\":\"" + abc + "\",\"def\":\"" + def + "\"}")
}
func (c *ChangeMsgController) Post() {
	abc := c.GetString("abc")
	def := c.GetString("def")
	fmt.Println(abc + ":" + def)
	c.Ctx.WriteString("{ \" abc\":\"" + abc + "\",\"def\":\"" + def + "\"}")
}

package controllers

import beego "github.com/beego/beego/v2/server/web"

type IframeController1 struct {
	beego.Controller
}
type IframeController2 struct {
	beego.Controller
}

func (c *IframeController1) Get() {
	c.TplName = "iframe1.html"

}
func (c *IframeController2) Get() {
	c.TplName = "iframe2.html"

}

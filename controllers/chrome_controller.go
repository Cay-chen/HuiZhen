package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// LoginController /*登录控制器*/
type ChromeController struct {
	beego.Controller
}

func (c *ChromeController) Get() {

	c.TplName = "chrome.html"

}

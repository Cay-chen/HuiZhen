package main

import (
	"HuiZhen/controllers"
	_ "HuiZhen/routers"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

func main() {

	// 注册自己写的错误页面
	beego.ErrorController(&controllers.ErrorController{})
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/`+time.Now().Format("2006-01-02")+`.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3000
	beego.Run()
}

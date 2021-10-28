package main

import (
	"HuiZhen/controllers"
	_ "HuiZhen/routers"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	// 注册自己写的错误页面
	beego.ErrorController(&controllers.ErrorController{})
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/hzxt.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	/*	data := make(map[string]string)
		data["JYConPersonPassword"] = "jy123456"
		data["method"] = "checkAccount"
		data["JYConPersonCode"] = "admin"
		data["JYConPersonBelongHos"] = "JYZX"
		b, _ := json.Marshal(data)*/
	//fmt.Println(b)
	//fmt.Println(utils.Post("http://222.209.81.188:9999/hzxt/JYConPersonServlet","method=checkAccount&JYConPersonCode=admin&JYConPersonPassword=jy123456&JYConPersonBelongHos=JYZX","application/x-www-form-urlencoded"))
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3000
	beego.Run()
	//fmt.Println(utils.Post("JYConPersonServlet","method=checkAccount&JYConPersonCode=admin&JYConPersonPassword=jy123456&JYConPersonBelongHos=JYZX"))

}

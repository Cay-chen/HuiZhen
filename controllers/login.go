package controllers

import (
	"HuiZhen/models/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

/*登录控制器*/
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		fmt.Println("notice:" + n)
		// 显示设置成功
		c.Redirect("/index", 302)
	} else if n, ok = flash.Data["error"]; ok {
		fmt.Println("Err:" + n)
		// 显示错误
		c.Data["Err"] = "inline"
		c.Data["Err_msg"] = n

		c.TplName = "login.html"
	} else {
		// 不然默认显示设置页面
		c.Data["Err"] = "none"
		c.TplName = "login.html"
	}
}

func (c *LoginController) Post() {

	serverName := "JYConPersonServlet"
	method := "checkAccount"
	JYConPersonCode := c.GetString("JYConPersonCode")
	JYConPersonPassword := c.GetString("JYConPersonPassword")
	JYConPersonBelongHos := c.GetString("JYConPersonBelongHos")
	parameterMap := make(map[string]string)
	parameterMap["JYConPersonCode"] = JYConPersonCode
	parameterMap["JYConPersonPassword"] = utils.Md5String32(JYConPersonPassword)
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	//parameter := "JYConPersonCode=" + JYConPersonCode + "&JYConPersonPassword=" + utils.Md5String32(JYConPersonPassword) + "&JYConPersonBelongHos=" + JYConPersonBelongHos
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05") +"  "+JYConPersonCode+" 登录账号： "+JYConPersonPassword)
	resmsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
	//resmsg :="{\"JYConPersonCode\": \"admin\",\"JYConPersonPassword\": \"c3469665bb653e9ff5ae280d5ed286b9\",\"JYConPersonBelongHos\": \"JYZX\",\"JYConPersonName\": \"管理员\",\"Flag\": \"true\",\"JYConPersonPhone\": \"18888888888\",\"JYConPersonType\": \"0\"}"
	if resmsg == "" {
		c.TplName = "400.html"
		return
	}
	fmt.Println("返回结果:" + resmsg)
	res := getPerson(resmsg)
	flash := beego.NewFlash()

	if res.Flag == "true" {
		err := c.SetSession("loginUser", resmsg)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  " + JYConPersonCode + " 登录成功" + JYConPersonPassword)
		if err != nil {
			return
		}
		flash.Notice(res.Mesg)
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	} else {
		if c.GetSession("loginUser") != nil {
			err := c.DestroySession()
			if err != nil {
				return
			}
		}
		flash.Error(res.Mesg)
		flash.Store(&c.Controller)

		c.Redirect("/", 302)

	}
}

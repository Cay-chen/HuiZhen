package controllers

import beego "github.com/beego/beego/v2/server/web"

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
	PersonUer *Person
	IsAdmin   bool
	IsYwb     bool
	IsKszr    bool
	IsYxys    bool
	IsFzr     bool
}

/**
判断是否登录
*/
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginUser")
	if loginuser != nil {
		this.IsLogin = true
		//this.LoginUser = loginuser
		this.PersonUer = getPerson(loginuser.(string))
		switch this.PersonUer.JYConPersonType {
		case "0":
			//jyConPersonType = "管理员"
			this.IsAdmin = true
			break
		case "1":
			//jyConPersonType = "科室主任"
			this.IsKszr = true

			break
		case "2":
			//jyConPersonType = "科室负责人"
			this.IsFzr = true
			break
		case "3":
			//jyConPersonType = "科室一线医生"
			this.IsYxys = true
			break
		case "4":
			//jyConPersonType = "医务部人员"
			this.IsYwb = true
			break
		}
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}

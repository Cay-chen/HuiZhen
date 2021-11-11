package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

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
type Person struct {
	JYConPersonType      string `json:"JYConPersonType"`
	JYConPersonBelongDep string `json:"JYConPersonBelongDep"`
	Flag                 string `json:"flag"`
	JYConPersonName      string `json:"JYConPersonName"`
	JYConPersonPhone     string `json:"JYConPersonPhone"`
	JYConPersonCode      string `json:"JYConPersonCode"`
	JYConPersonBelongHos string `json:"JYConPersonBelongHos"`
	Mesg                 string `json:"mesg"`
}

/**
判断是否登录
*/
func (this *BaseController) Prepare() {
	loginUser := this.GetSession("loginUser")
	if loginUser != nil {
		this.IsLogin = true
		//this.LoginUser = loginuser
		this.PersonUer = getPerson(loginUser.(string))
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
func getPerson(msg string) (person *Person) {

	res := &Person{}
	err := json.Unmarshal([]byte(msg), &res)
	if err != nil {
		return
	}
	return res
}

type Msg struct {
	Flag string `json:"flag"`
	Mesg string `json:"mesg"`
}

func getMsg(msg string) (person *Msg) {

	res := &Msg{}
	err := json.Unmarshal([]byte(msg), &res)
	if err != nil {
		return
	}
	return res
}

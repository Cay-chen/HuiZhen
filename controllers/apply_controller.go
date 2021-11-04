package controllers

import (
	"HuiZhen/models"
	"HuiZhen/models/utils"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ApplyController struct {
	BaseController
}

func (c *ApplyController) Get() {
	if c.IsLogin {
		c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200").Data
		flash := beego.ReadFromRequest(&c.Controller)
		if n, ok := flash.Data["notice"]; ok {
			fmt.Println("notice:" + n)
			c.Data["Msg_ture"] = true
			// 显示设置成功
			c.Data["Msg_notice"] = n
		} else if n, ok = flash.Data["error"]; ok {
			c.Data["Msg_ture"] = true
			// 显示错误
			c.Data["Msg_notice"] = n
		} else {
			// 不然默认显示设置页面
			c.Data["Msg_notice"] = ""
		}
		editInfo := false
		JYConNum := c.GetString("JYConNum")
		ConType1 := "checked"
		ConType2 := " "
		selectedW := " "
		selectedM := "selected"
		if JYConNum != "" {
			info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "")
			s := models.FormInfo{}
			err := json.Unmarshal([]byte(info), &s)
			if err != nil {
			}
			c.Data["JYConSickName"] = s.JYConSickName
			c.Data["JYConSickSex"] = s.JYConSickSex

			/*if s.JYConSickSex =="男"{
				selectedM="selected"
			}*/
			if s.JYConSickSex == "女" {
				selectedM = " "
				selectedW = "selected"
			}
			editInfo = true
			c.Data["SexW"] = selectedW
			c.Data["JYConNum"] = JYConNum
			c.Data["SexM"] = selectedM
			c.Data["JYConSickAge"] = s.JYConSickAge
			c.Data["JYConSickDep"] = s.JYConSickDep
			c.Data["JYConSickBedNo"] = s.JYConSickBedNo
			c.Data["JYConSickAd"] = s.JYConSickAd
			c.Data["JYConDepLocaltion"] = s.JYConDepLocaltion
			c.Data["JYConSickDoc"] = s.JYConSickDoc
			c.Data["JYConSickDocPhone"] = s.JYConSickDocPhone
			c.Data["JYConType"] = s.JYConType
			/*if s.JYConType == "1" {
				ConType1 = "checked"
			}*/
			if s.JYConType == "2" {
				ConType1 = " "
				ConType2 = "checked"
			}
			c.Data["JYConSickDia"] = s.JYConSickDia
			c.Data["JYConSickCase"] = s.JYConSickCase
			c.Data["JYConPurpose"] = s.JYConPurpose
			c.Data["JYConOppHos"] = s.JYConOppHos
			c.Data["JYConOppDepId"] = s.JYConOppDepId
			c.Data["JYConDate"] = s.JYConDate
		}
		c.Data["SexW"] = selectedW
		c.Data["SexM"] = selectedM
		c.Data["ConType2"] = ConType2
		c.Data["ConType1"] = ConType1
		c.Data["EditInfo"] = editInfo
		c.TplName = "apply.html"
	} else {
		c.Redirect("/error/56", 302)
		//	c.TplName = "error.html"
	}
}
func (c *ApplyController) Post() {
	if c.IsLogin {
		JYConNum := c.GetString("JYConNum")
		JYConSickName := c.GetString("JYConSickName")
		JYConSickSex := c.GetString("JYConSickSex")
		JYConSickAge := c.GetString("JYConSickAge")
		JYConSickDepId := c.GetString("JYConSickDepId")
		JYConSickDep := c.GetString("JYConSickDep")
		JYConSickBedNo := c.GetString("JYConSickBedNo")
		JYConSickAd := c.GetString("JYConSickAd")
		JYConDepLocaltion := c.GetString("JYConDepLocaltion")
		JYConSickDocId := c.GetString("JYConSickDocId")
		JYConSickDoc := c.GetString("JYConSickDoc")
		JYConSickDocPhone := c.GetString("JYConSickDocPhone")
		JYConType := c.GetString("JYConType")
		JYConSickDia := c.GetString("JYConSickDia")
		JYConSickCase := c.GetString("JYConSickCase")
		JYConPurpose := c.GetString("JYConPurpose")
		JYConOppDep := c.GetString("JYConOppDep")
		JYConOppHos := c.GetString("JYConOppHos")
		flag := c.GetString("flag")
		JYConOppDepId := c.GetString("JYConOppDepId")
		JYConDate := c.GetString("JYConDate")
		//fmt.Println("JYConSickName:"+JYConSickName+" | JYConSickSex:"+JYConSickSex+" | JYConSickAge:"+JYConSickAge+" | JYConSickDepId:"+JYConSickDepId+" | JYConSickDep:"+JYConSickDep+" | JYConSickBelongHos:"+JYConSickBelongHos+" | JYConSickBedNo:"+JYConSickBedNo+" | JYConSickAd:"+JYConSickAd+" | JYConDepLocaltion:"+JYConDepLocaltion+" | JYConSickDocId:"+JYConSickDocId+" | JYConSickDoc:"+JYConSickDoc+" | JYConSickDocPhone:"+JYConSickDocPhone+" | JYConType:"+JYConType+" | JYConSickDia:"+JYConSickDia+" | JYConSickCase:"+JYConSickCase+" | JYConPurpose:"+JYConPurpose+" | JYConOppDep:"+JYConOppDep+" | JYConOppDepId:"+JYConOppDepId+" | JYConDate:"+JYConDate+" | JYConFormCreatePersonId:"+JYConFormCreatePersonId+" | JYConFormCreatePersonName:"+JYConFormCreatePersonName)

		parameterMap := make(map[string]string)
		parameterMap["JYConSickName"] = JYConSickName
		parameterMap["JYConSickSex"] = JYConSickSex
		parameterMap["JYConSickAge"] = JYConSickAge //性别
		parameterMap["JYConSickDepId"] = JYConSickDepId
		parameterMap["JYConSickDep"] = JYConSickDep //科室名称
		parameterMap["JYConSickBelongHos"] = c.PersonUer.JYConPersonBelongHos
		parameterMap["JYConSickBedNo"] = JYConSickBedNo
		parameterMap["JYConSickAd"] = JYConSickAd
		parameterMap["JYConDepLocaltion"] = JYConDepLocaltion
		parameterMap["JYConSickDocId"] = JYConSickDocId
		parameterMap["JYConSickDoc"] = JYConSickDoc //医生姓名
		parameterMap["JYConSickDocPhone"] = JYConSickDocPhone
		parameterMap["JYConType"] = JYConType
		parameterMap["JYConSickDia"] = JYConSickDia
		parameterMap["JYConSickCase"] = JYConSickCase
		parameterMap["JYConPurpose"] = JYConPurpose
		parameterMap["JYConOppDep"] = JYConOppDep
		parameterMap["JYConOppDepId"] = JYConOppDepId
		parameterMap["flag"] = flag
		parameterMap["JYConOppHos"] = JYConOppHos
		parameterMap["JYConDate"] = JYConDate
		parameterMap["JYConFormCreatePersonId"] = c.PersonUer.JYConPersonCode
		parameterMap["JYConFormCreatePersonName"] = c.PersonUer.JYConPersonName
		postResult := ""
		if JYConNum != "" {
			serverName := "JYConFormServlet"
			method := "editForm"
			parameterMap["JYConNum"] = JYConNum
			postResult = utils.Post(serverName, method, utils.MapToUrl(parameterMap))
		} else {
			serverName := "JYConFormServlet"
			method := "createForm"
			postResult = utils.Post(serverName, method, utils.MapToUrl(parameterMap))
		}

		res := getMsg(postResult)
		//res:="true"
		flash := beego.NewFlash()
		if res.Flag == "true" {
			flash.Notice(res.Mesg)
			//flash.Notice("申请成功")
		} else {
			flash.Error(res.Mesg)
			//flash.Error("申请失败")
		}
		flash.Store(&c.Controller)
		c.Redirect("/apply", 302)

	}

}
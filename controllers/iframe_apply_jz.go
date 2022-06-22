package controllers

import (
	"HuiZhen/models"
	"HuiZhen/models/utils"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type ApplicationFormEtController struct {
	BaseController
}

func (c *ApplicationFormEtController) Get() {
	if c.IsLogin {
		depList := models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200", c.PersonUer.JYConPersonCode).Data
		//c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200").Data
		flash := beego.ReadFromRequest(&c.Controller)
		if n, ok := flash.Data["notice"]; ok {
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
		JYConOppHos := ""
		DepListS := ""
		if c.PersonUer.JYConPersonBelongHos == "JYZX" {
			c.Data["IsJY"] = true
		} else {
			c.Data["IsJY"] = false
		}
		Y1 := false
		for i := 0; i < len(depList); i++ {
			DepListS = DepListS + "                    <option value=\"" + depList[i].JYConDepCode + "\" >" + depList[i].JYConDepName + "</option>\n"
		}
		if JYConNum != "" {
			DepListS = ""
			info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "", c.PersonUer.JYConPersonCode)
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
			for i := 0; i < len(depList); i++ {
				if s.JYConSickDepId == depList[i].JYConDepCode {
					DepListS = DepListS + "                    <option value=\"" + depList[i].JYConDepCode + "\" selected>" + depList[i].JYConDepName + "</option>\n"
				} else {
					DepListS = DepListS + "                    <option value=\"" + depList[i].JYConDepCode + "\" >" + depList[i].JYConDepName + "</option>\n"
				}
			}
			editInfo = true
			Y1 = true
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
			c.Data["JYConFormConclusion"] = s.JYConFormConclusion
			c.Data["JYConFormConclusionDate"] = s.JYConFormConclusionDate
			c.Data["JYConType"] = s.JYConType
			/*if s.JYConType == "1" {
				ConType1 = "checked"
			}*/
			if s.JYConType == "2" {
				ConType1 = " "
				ConType2 = "checked"
			}
			c.Data["JYConSickDia"] = utils.StringZyt(s.JYConSickDia)
			c.Data["JYConSickDepId"] = s.JYConSickDepId
			c.Data["JYConSickCase"] = utils.StringZyt(s.JYConSickCase)
			c.Data["JYConPurpose"] = utils.StringZyt(s.JYConPurpose)
			c.Data["JYConSickDocId"] = s.JYConSickDocId
			c.Data["JYConOppDocPhone"] = s.JYConOppDocPhone
			c.Data["JYConOppDocName"] = s.JYConOppDocName
			JYConOppHos = s.JYConOppHos
			c.Data["JYConOppDepId"] = s.JYConOppDepId
			c.Data["JYConDate"] = string([]byte(s.JYConDate)[:19])
			c.Data["JYConOppDep"] = s.JYConOppDep
			c.Data["Yq"] = s.JYConOppHos
			c.Data["JYConOppDocId"] = s.JYConOppDocId
		}
		c.Data["JYConOppHos"] = JYConOppHos
		c.Data["SexW"] = selectedW
		c.Data["SexM"] = selectedM
		c.Data["ConType2"] = ConType2
		c.Data["ConType1"] = ConType1
		c.Data["EditInfo"] = editInfo
		c.Data["Y1"] = Y1
		c.Data["DepListS"] = DepListS
		c.TplName = "iframe_apply_jz.html"
	} else {
		c.Redirect("/error/600", 302)
	}
}
func (c *ApplicationFormEtController) Post() {
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
		JYConDate := c.GetString("JYConDate")
		JYConOppDepId := c.GetString("JYConOppDepId")
		JYConOppDocName := c.GetString("JYConOppDocName")
		JYConOppDocId := c.GetString("JYConOppDocId")
		JYConOppDocPhone := c.GetString("JYConOppDocPhone")
		JYConFormConclusion := c.GetString("JYConFormConclusion")
		JYConFormConclusionDate := c.GetString("JYConFormConclusionDate")
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
		//parameterMap["JYConSickDia"] = JYConSickDia
		//parameterMap["JYConSickCase"] = JYConSickCase
		//parameterMap["JYConPurpose"] = JYConPurpose
		parameterMap["JYConSickDia"] = utils.StringZy(JYConSickDia)
		parameterMap["JYConSickCase"] = utils.StringZy(JYConSickCase)
		parameterMap["JYConPurpose"] = utils.StringZy(JYConPurpose)
		parameterMap["JYConOppDep"] = JYConOppDep
		parameterMap["JYConOppDepId"] = JYConOppDepId
		parameterMap["flag"] = flag
		parameterMap["JYConOppHos"] = JYConOppHos
		parameterMap["JYConDate"] = JYConDate
		parameterMap["JYConOppDocName"] = JYConOppDocName
		parameterMap["JYConOppDocId"] = JYConOppDocId
		parameterMap["JYConOppDocPhone"] = JYConOppDocPhone
		parameterMap["JYConFormConclusion"] = JYConFormConclusion
		parameterMap["JYConFormConclusionDate"] = JYConFormConclusionDate
		parameterMap["JYConFormCreatePersonId"] = c.PersonUer.JYConPersonCode
		parameterMap["JYConFormCreatePersonName"] = c.PersonUer.JYConPersonName
		postResult := ""
		if JYConNum != "" {
			serverName := "JYConFormUrgentServlet"
			method := "editFormUrgent"
			parameterMap["JYConNum"] = JYConNum
			postResult = utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
		} else {
			serverName := "JYConFormUrgentServlet"
			method := "createUrgentForm"
			postResult = utils.Post(serverName, method, utils.MapToUrl(parameterMap), c.PersonUer.JYConPersonCode)
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
		c.Redirect("/apply_et", 302)

	}

}

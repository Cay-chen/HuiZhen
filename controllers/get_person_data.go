package controllers

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type GetPersonData struct {
	Code int `json:"code"`
	Data []struct {
		JYConPersonType      string `json:"JYConPersonType"`
		JYConPersonBelongDep string `json:"JYConPersonBelongDep"`
		JYConPersonPassword  string `json:"JYConPersonPassword"`
		JYConPersonName      string `json:"JYConPersonName"`
		JYConPersonPhone     string `json:"JYConPersonPhone"`
		JYConPersonIsActive  string `json:"JYConPersonIsActive"`
		JYConPersonAutoNum   int    `json:"JYConPersonAutoNum"`
		JYConPersonCode      string `json:"JYConPersonCode"`
		JYConPersonBelongHos string `json:"JYConPersonBelongHos"`
		JYConDepName         string `json:"JYConDepName"`
	} `json:"data"`
	Count int    `json:"count"`
	Msg   string `json:"msg"`
}

type GetPersonData1 struct {
	Code string `json:"code"`
	Data []struct {
		JYConPersonType      string `json:"JYConPersonType"`
		JYConPersonBelongDep string `json:"JYConPersonBelongDep"`
		JYConPersonPassword  string `json:"JYConPersonPassword"`
		JYConPersonName      string `json:"JYConPersonName"`
		JYConPersonPhone     string `json:"JYConPersonPhone"`
		JYConPersonIsActive  string `json:"JYConPersonIsActive"`
		JYConPersonAutoNum   int    `json:"JYConPersonAutoNum"`
		JYConPersonCode      string `json:"JYConPersonCode"`
		JYConPersonBelongHos string `json:"JYConPersonBelongHos"`
		JYConDepName         string `json:"JYConDepName"`
	} `json:"data"`
	Count string `json:"count"`
	Mesg  string `json:"mesg"`
}

//设置返回类型为string
func Data1ToDataJson(c GetPersonData1) string {
	var personData GetPersonData
	//将结构体中的Name的值变为字符串1
	code, _ := strconv.Atoi(c.Code)
	personData.Code = code
	count, _ := strconv.Atoi(c.Count)
	personData.Count = count
	//personData.Count = 30
	personData.Data = c.Data
	personData.Msg = c.Mesg
	//人员类型（0--管理员，1--科室主任，2--科室负责人，3--科室一线医生,4-医务部人员）
	for i := 0; i < len(personData.Data); i++ {
		jyConPersonType := personData.Data[i].JYConPersonType
		switch jyConPersonType {
		case "0":
			jyConPersonType = "管理员"
			break
		case "1":
			jyConPersonType = "科室主任"
			break
		case "2":
			jyConPersonType = "科室负责人"
			break
		case "3":
			jyConPersonType = "科室一线医生"
			break
		case "4":
			jyConPersonType = "医务部人员"
			break
		}
		personData.Data[i].JYConPersonType = jyConPersonType

		jyConPersonIsActive := personData.Data[i].JYConPersonIsActive
		switch jyConPersonIsActive {
		case "0":
			jyConPersonIsActive = "未激活"
			break
		case "1":
			jyConPersonIsActive = "已激活"
			break
		}
		personData.Data[i].JYConPersonIsActive = jyConPersonIsActive

	}
	body, _ := json.Marshal(personData)
	return string(body)
}

type GetPersonDataController struct {
	BaseController
}

func (c *GetPersonDataController) Get() {
	if c.IsLogin {
		if c.IsAdmin || c.IsYwb {
			limit := c.GetString("limit")
			page := c.GetString("page")
			JYConPersonBelongHos := c.PersonUer.JYConPersonBelongHos
			JYConDepCode := c.GetString("key[JYConDepCode]")
			JYConPersonCode := c.GetString("key[JYConPersonCode]")
			serverName := "JYConPersonServlet"
			method := "getPersonInfoByHos"
			parameterMap := make(map[string]string)
			parameterMap["limit"] = limit
			parameterMap["page"] = page
			parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
			if JYConDepCode != "" {
				parameterMap["JYConDepCode"] = JYConDepCode
			}
			if JYConPersonCode != "" {
				parameterMap["JYConPersonCode"] = JYConPersonCode
			}
			postResult := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			if postResult != "" {
				fmt.Println("POST返回结果:" + postResult)
				res := GetPersonData1{}
				_ = json.Unmarshal([]byte(postResult), &res)
				//Data1ToDataJson(res)
				fmt.Println("查询结果：" + Data1ToDataJson(res))
				c.Ctx.WriteString(Data1ToDataJson(res))
			} else {
				c.TplName = "error.html"
				//处理数据为空
			}
		} else {
			c.Redirect("/error/5434", 302)
		}

	} else {
		c.Redirect("/error/5434", 302)
	}

}

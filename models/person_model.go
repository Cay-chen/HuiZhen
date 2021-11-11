package models

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"fmt"
	"strconv"
)

type Person struct {
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
type Person1 struct {
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

func GetDocList(JYConPersonBelongHos, JYConDepCode, JYConPersonType, limit, page string) string {
	serverName := "JYConPersonServlet"
	method := "getPersonInfoByHos"
	parameterMap := make(map[string]string)
	parameterMap["limit"] = limit
	parameterMap["page"] = page
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	if JYConDepCode != "" {
		parameterMap["JYConPersonBelongDep"] = JYConDepCode
	}
	if JYConPersonType != "" {
		parameterMap["JYConPersonType"] = JYConPersonType
	}
	postResult := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
	if postResult != "" {
		fmt.Println("POST返回结果:" + postResult)
		res := Person1{}
		_ = json.Unmarshal([]byte(postResult), &res)
		//Data1ToDataJson(res)
		fmt.Println("查询结果：" + data1ToDataJson(res))
		return data1ToDataJson(res)

	} else {
		return ""
	}
}
func GetDocInfo(JYConPersonCode, JYConPersonBelongHos string) string {
	serverName := "JYConPersonServlet"
	method := "getPersonInfoByAccount"
	parameterMap := make(map[string]string)
	parameterMap["JYConPersonCode"] = JYConPersonCode
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
	return resMsg

}

//设置返回类型为string
func data1ToDataJson(c Person1) string {
	var personData Person
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

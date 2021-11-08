package models

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"strconv"
)

type ExtraInfoJson struct {
	Code  int         `json:"code"`
	Data  []ExtraInfo `json:"data"`
	Count int         `json:"count"`
	Msg   string      `json:"mesg"`
}

type ExtraInfoJson1 struct {
	Code  string      `json:"code"`
	Data  []ExtraInfo `json:"data"`
	Count string      `json:"count"`
	Mesg  string      `json:"mesg"`
}

type ExtraInfo struct {
	JYConPersonType      string      `json:"JYConPersonType"`
	JYConPersonBelongDep string      `json:"JYConPersonBelongDep"`
	JYConPersonName      string      `json:"JYConPersonName"`
	JYConDepName         string      `json:"JYConDepName"`
	JYConPersonPhone     string      `json:"JYConPersonPhone"`
	JYConPersonFromNum   string      `json:"JYConPersonFromNum"`
	JYConPersonAutoNum   int         `json:"JYConPersonAutoNum"`
	JYConPersonState     string      `json:"JYConPersonState"`
	JYConPersonCode      interface{} `json:"JYConPersonCode"`
	JYConPersonBelongHos string      `json:"JYConPersonBelongHos"`
}

func extraInfoDate1ToData(c ExtraInfoJson1) ExtraInfoJson {
	var depInfo ExtraInfoJson
	//将结构体中的Name的值变为字符串1
	code, _ := strconv.Atoi(c.Code)
	depInfo.Code = code
	count, _ := strconv.Atoi(c.Count)
	depInfo.Count = count
	//personData.Count = 30
	depInfo.Data = c.Data
	depInfo.Msg = c.Mesg
	return depInfo
}

func GetExtraInfoList(JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit string) ExtraInfoJson {
	serverName := "JYConPersonExtraInfoServlet"
	method := "queryExtraInfo"
	parameterMap := make(map[string]string)
	parameterMap["JYConPersonType"] = JYConPersonType
	parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
	parameterMap["JYConPersonCode"] = JYConPersonCode
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	parameterMap["page"] = page
	parameterMap["limit"] = limit
	resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
	res1 := ExtraInfoJson1{}
	_ = json.Unmarshal([]byte(resMsg), &res1)
	res := extraInfoDate1ToData(res1)
	return res
}

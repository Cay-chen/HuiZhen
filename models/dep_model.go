package models

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"strconv"
)

type DepInfo struct {
	Code  int    `json:"code"`
	Data  []Dep  `json:"data"`
	Count int    `json:"count"`
	Msg   string `json:"mesg"`
}
type DepInfo1 struct {
	Code  string `json:"code"`
	Data  []Dep  `json:"data"`
	Count string `json:"count"`
	Mesg  string `json:"mesg"`
}
type Dep struct {
	JYConDepLocalhost string `json:"JYConDepLocalhost"`
	JYDepAutoNum      int    `json:"JYDepAutoNum"`
	JYConDepCode      string `json:"JYConDepCode"`
	JYConDepName      string `json:"JYConDepName"`
	JYConDepBelongHos string `json:"JYConDepBelongHos"`
	JYConDepPhone     string `json:"JYConDepPhone"`
	JYConDepIsActive  string `json:"JYConDepIsActive"`
}

func GetDepList(JYConPersonBelongHos, page, limit, postName string) DepInfo {

	serverName := "JYConDepListServlet"
	method := "getDepListByHos"
	parameterMap := make(map[string]string)
	parameterMap["JYConDepBelongHos"] = JYConPersonBelongHos
	parameterMap["page"] = page
	parameterMap["limit"] = limit
	resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), postName)
	res1 := DepInfo1{}
	_ = json.Unmarshal([]byte(resMsg), &res1)
	res := data1ToData(res1)
	//_ = json.Unmarshal([]byte(resMsg), &res)
	return res
}
func data1ToData(c DepInfo1) DepInfo {
	var depInfo DepInfo
	//将结构体中的Name的值变为字符串1
	code, _ := strconv.Atoi(c.Code)
	depInfo.Code = code
	count, _ := strconv.Atoi(c.Count)
	depInfo.Count = count
	//personData.Count = 30
	depInfo.Data = c.Data
	depInfo.Msg = c.Mesg
	//人员类型（0--管理员，1--科室主任，2--科室负责人，3--科室一线医生,4-医务部人员）
	for i := 0; i < len(depInfo.Data); i++ {
		jyConPersonIsActive := depInfo.Data[i].JYConDepIsActive
		switch jyConPersonIsActive {
		case "0":
			jyConPersonIsActive = "未激活"
			break
		case "1":
			jyConPersonIsActive = "已激活"
			break
		}
		depInfo.Data[i].JYConDepIsActive = jyConPersonIsActive

	}
	return depInfo
}

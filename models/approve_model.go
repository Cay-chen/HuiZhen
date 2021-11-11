package models

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"strconv"
)

type ApproveJson struct {
	Code  int       `json:"code"`
	Data  []Approve `json:"data"`
	Count int       `json:"count"`
	Msg   string    `json:"mesg"`
}

type ApproveJson1 struct {
	Code  string    `json:"code"`
	Data  []Approve `json:"data"`
	Count string    `json:"count"`
	Mesg  string    `json:"mesg"`
}

type Approve struct {
	JYConFormApproveOldState   string `json:"JYConFormApproveOldState"`
	JYConFormApproveState      string `json:"JYConFormApproveState"`
	JYConFormApproveDate       string `json:"JYConFormApproveDate"`
	JYConFormApproveBelongHos  string `json:"JYConFormApproveBelongHos"`
	JYFormappAutoNum           int    `json:"JYFormappAutoNum"`
	JYConNum                   string `json:"JYConNum"`
	JYConFormApproveNewState   string `json:"JYConFormApproveNewState"`
	JYConFormApprovePersonName string `json:"JYConFormApprovePersonName"`
	JYConFormApproveComment    string `json:"JYConFormApproveComment"`
	JYConFormApprovePersonId   string `json:"JYConFormApprovePersonId"`
}

func approveDate1ToData(c ApproveJson1) ApproveJson {
	var depInfo ApproveJson
	//将结构体中的Name的值变为字符串1
	code, _ := strconv.Atoi(c.Code)
	depInfo.Code = code
	count, _ := strconv.Atoi(c.Count)
	depInfo.Count = count
	//personData.Count = 30
	depInfo.Data = c.Data
	depInfo.Msg = c.Mesg
	//更改时间
	for i := 0; i < len(depInfo.Data); i++ {
		JYConFormCreateDate := depInfo.Data[i].JYConFormApproveDate
		data := string([]byte(JYConFormCreateDate)[:19])
		depInfo.Data[i].JYConFormApproveDate = data
	}
	return depInfo
}
func GetPersonApproveList(JYConPersonCode, JYConPersonBelongHos, flag, limit, page, postName string) ApproveJson {
	serverName := "JYConFormServlet"
	method := "getMyApprovalFormEd"
	parameterMap := make(map[string]string)
	parameterMap["JYConPersonCode"] = JYConPersonCode
	parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
	parameterMap["flag"] = flag
	parameterMap["limit"] = limit
	parameterMap["page"] = page
	resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), postName)
	res1 := ApproveJson1{}
	_ = json.Unmarshal([]byte(resMsg), &res1)
	res := approveDate1ToData(res1)
	return res
}
func GetApproveList(JYConNum, limit, page, postName string) ApproveJson {
	serverName := "JYConFormApproveServlet"
	method := "getApproveInfo"
	parameterMap := make(map[string]string)
	parameterMap["JYConNum"] = JYConNum
	parameterMap["limit"] = limit
	parameterMap["page"] = page
	resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap), postName)
	res1 := ApproveJson1{}
	_ = json.Unmarshal([]byte(resMsg), &res1)
	res := approveDate1ToData(res1)
	return res
}

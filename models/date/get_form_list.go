package date

import (
	"HuiZhen/models/utils"
	"encoding/json"
	"strconv"
)

type FormInfo struct {
	Code int `json:"code"`
	Data []struct {
		JYConSickDocPhone         string `json:"JYConSickDocPhone"`
		JYConSickDepId            string `json:"JYConSickDepId"`
		JYConSickCase             string `json:"JYConSickCase"`
		JYConOppDocPhone          string `json:"JYConOppDocPhone"`
		JYConSickDep              string `json:"JYConSickDep"`
		JYConFormCreateDate       string `json:"JYConFormCreateDate"`
		JYConSickBedNo            string `json:"JYConSickBedNo"`
		JYConSickDia              string `json:"JYConSickDia"`
		JYConOppDocId             string `json:"JYConOppDocId"`
		JYConNum                  string `json:"JYConNum"`
		JYConSickAge              string `json:"JYConSickAge"`
		JYConFormPolicy           string `json:"JYConFormPolicy"`
		JYConSickName             string `json:"JYConSickName"`
		JYConFormModifyDate       string `json:"JYConFormModifyDate"`
		JYConSickDocId            string `json:"JYConSickDocId"`
		JYConSickDoc              string `json:"JYConSickDoc"`
		JYConSickAd               string `json:"JYConSickAd"`
		JYConOppDocName           string `json:"JYConOppDocName"`
		JYConSickBelongHos        string `json:"JYConSickBelongHos"`
		JYConDepLocaltion         string `json:"JYConDepLocaltion"`
		JYConOppHos               string `json:"JYConOppHos"`
		JYConDate                 string `json:"JYConDate"`
		JYConSickSex              string `json:"JYConSickSex"`
		JYConOppDepId             string `json:"JYConOppDepId"`
		JYConFormCreatePersonName string `json:"JYConFormCreatePersonName"`
		JYConOppDep               string `json:"JYConOppDep"`
		JYConFormCreatePersonId   string `json:"JYConFormCreatePersonId"`
		JYConPurpose              string `json:"JYConPurpose"`
		JYConType                 string `json:"JYConType"`
	} `json:"data"`
	Count int    `json:"count"`
	Msg   string `json:"mesg"`
}
type FormInfo1 struct {
	Code string `json:"code"`
	Data []struct {
		JYConSickDocPhone         string `json:"JYConSickDocPhone"`
		JYConSickDepId            string `json:"JYConSickDepId"`
		JYConSickCase             string `json:"JYConSickCase"`
		JYConOppDocPhone          string `json:"JYConOppDocPhone"`
		JYConSickDep              string `json:"JYConSickDep"`
		JYConFormCreateDate       string `json:"JYConFormCreateDate"`
		JYConSickBedNo            string `json:"JYConSickBedNo"`
		JYConSickDia              string `json:"JYConSickDia"`
		JYConOppDocId             string `json:"JYConOppDocId"`
		JYConNum                  string `json:"JYConNum"`
		JYConSickAge              string `json:"JYConSickAge"`
		JYConFormPolicy           string `json:"JYConFormPolicy"`
		JYConSickName             string `json:"JYConSickName"`
		JYConFormModifyDate       string `json:"JYConFormModifyDate"`
		JYConSickDocId            string `json:"JYConSickDocId"`
		JYConSickDoc              string `json:"JYConSickDoc"`
		JYConSickAd               string `json:"JYConSickAd"`
		JYConOppDocName           string `json:"JYConOppDocName"`
		JYConSickBelongHos        string `json:"JYConSickBelongHos"`
		JYConDepLocaltion         string `json:"JYConDepLocaltion"`
		JYConOppHos               string `json:"JYConOppHos"`
		JYConDate                 string `json:"JYConDate"`
		JYConSickSex              string `json:"JYConSickSex"`
		JYConOppDepId             string `json:"JYConOppDepId"`
		JYConFormCreatePersonName string `json:"JYConFormCreatePersonName"`
		JYConOppDep               string `json:"JYConOppDep"`
		JYConFormCreatePersonId   string `json:"JYConFormCreatePersonId"`
		JYConPurpose              string `json:"JYConPurpose"`
		JYConType                 string `json:"JYConType"`
	} `json:"data"`
	Count string `json:"count"`
	Mesg  string `json:"mesg"`
}
type FormInfo2 struct {
	JYConSickDocPhone         string `json:"JYConSickDocPhone"`
	JYConSickDepId            string `json:"JYConSickDepId"`
	JYConSickCase             string `json:"JYConSickCase"`
	JYConOppDocPhone          string `json:"JYConOppDocPhone"`
	JYConSickDep              string `json:"JYConSickDep"`
	JYConFormCreateDate       string `json:"JYConFormCreateDate"`
	JYConSickBedNo            string `json:"JYConSickBedNo"`
	JYConSickDia              string `json:"JYConSickDia"`
	JYConOppDocId             string `json:"JYConOppDocId"`
	JYConNum                  string `json:"JYConNum"`
	JYConSickAge              string `json:"JYConSickAge"`
	JYConFormPolicy           string `json:"JYConFormPolicy"`
	JYConSickName             string `json:"JYConSickName"`
	JYConFormModifyDate       string `json:"JYConFormModifyDate"`
	JYConSickDocId            string `json:"JYConSickDocId"`
	JYConSickDoc              string `json:"JYConSickDoc"`
	JYConSickAd               string `json:"JYConSickAd"`
	JYConOppDocName           string `json:"JYConOppDocName"`
	JYConSickBelongHos        string `json:"JYConSickBelongHos"`
	JYConDepLocaltion         string `json:"JYConDepLocaltion"`
	JYConOppHos               string `json:"JYConOppHos"`
	JYConDate                 string `json:"JYConDate"`
	JYConSickSex              string `json:"JYConSickSex"`
	JYConOppDepId             string `json:"JYConOppDepId"`
	JYConFormCreatePersonName string `json:"JYConFormCreatePersonName"`
	JYConOppDep               string `json:"JYConOppDep"`
	JYConFormCreatePersonId   string `json:"JYConFormCreatePersonId"`
	JYConPurpose              string `json:"JYConPurpose"`
	JYConType                 string `json:"JYConType"`
}

func GetFormList(flag, JYConPersonType, JYConPersonBelongDep, JYConPersonCode, JYConPersonBelongHos, page, limit, JYConNum, methods, flagForm string) string {
	serverName := "JYConFormServlet"
	backRes := ""
	if methods == "getFormBySomething" {
		method := "getFormBySomething"
		parameterMap := make(map[string]string)
		if flag == "one" {
			parameterMap["flag"] = flag
			parameterMap["JYConNum"] = JYConNum
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			backRes = resMsg

		}
		if flag == "many" {
			parameterMap["flag"] = flag
			parameterMap["JYConPersonType"] = JYConPersonType
			parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
			parameterMap["JYConPersonCode"] = JYConPersonCode
			parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
			parameterMap["page"] = page
			parameterMap["limit"] = limit
			resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
			res1 := FormInfo1{}
			_ = json.Unmarshal([]byte(resMsg), &res1)
			res := formDate1ToData(res1)
			//_ = json.Unmarshal([]byte(resMsg), &res)
			body, _ := json.Marshal(res)
			backRes = string(body)
		}

	} else if methods == "getMyApprovalForm" {
		method := "getMyApprovalForm"
		parameterMap := make(map[string]string)
		parameterMap["flag"] = flagForm
		parameterMap["JYConPersonType"] = JYConPersonType
		parameterMap["JYConPersonBelongDep"] = JYConPersonBelongDep
		parameterMap["JYConPersonCode"] = JYConPersonCode
		parameterMap["JYConPersonBelongHos"] = JYConPersonBelongHos
		parameterMap["page"] = page
		parameterMap["limit"] = limit
		resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
		res1 := FormInfo1{}
		_ = json.Unmarshal([]byte(resMsg), &res1)
		res := formDate1ToData(res1)
		//_ = json.Unmarshal([]byte(resMsg), &res)
		body, _ := json.Marshal(res)
		backRes = string(body)

	}

	/*resMsg := utils.Post(serverName, method, utils.MapToUrl(parameterMap))
	res1 := FormInfo1{}
	_ = json.Unmarshal([]byte(resMsg), &res1)
	res:=formDate1ToData(res1)
	//_ = json.Unmarshal([]byte(resMsg), &res)
	body, _ := json.Marshal(res)
	return string(body)*/
	return backRes
}
func formDate1ToData(c FormInfo1) FormInfo {
	var depInfo FormInfo
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
		jYConType := depInfo.Data[i].JYConType
		switch jYConType {
		case "1":
			jYConType = "平会诊"
			break
		case "2":
			jYConType = "急会诊"
			break
		}
		depInfo.Data[i].JYConType = jYConType

	}
	return depInfo
}

package models

import (
	"HuiZhen/models/utils"
	"encoding/json"
)

type HisInfo struct {
	SexCode      string `json:"SEX_CODE"`
	DeptName     string `json:"DEPT_NAME"`
	BedNO        string `json:"BED_NO"`
	PatientNO    string `json:"PATIENT_NO"`
	DeptCode     string `json:"DEPT_CODE"`
	MainDiagnose string `json:"MAIN_DIAGNOSE"`
	Name         string `json:"NAME"`
	Age          int    `json:"AGE"`
}

func GetHisInfo(JYConSickAd, JYConSickBelongHos, postName string) HisInfo {
	serverName := "JYConSickInfoFromHisServlet"
	method := "querySickInfoFromHis"
	paramMap := make(map[string]string)
	paramMap["JYConSickAd"] = JYConSickAd
	paramMap["JYConSickBelongHos"] = JYConSickBelongHos
	resMsg := utils.Post(serverName, method, utils.MapToUrl(paramMap), postName)
	res := HisInfo{}
	_ = json.Unmarshal([]byte(resMsg), &res)
	return res
}

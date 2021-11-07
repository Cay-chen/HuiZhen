package models

import "strconv"

type JsonInt struct {
	Code  int           `json:"code"`
	Data  []interface{} `json:"data"`
	Count int           `json:"count"`
	Msg   string        `json:"mesg"`
}
type JsonString struct {
	Code  string        `json:"code"`
	Data  []interface{} `json:"data"`
	Count string        `json:"count"`
	Mesg  string        `json:"mesg"`
}
type ExtraInfo1 struct {
	JYConPersonType      string      `json:"JYConPersonType"`
	JYConPersonBelongDep string      `json:"JYConPersonBelongDep"`
	JYConPersonName      string      `json:"JYConPersonName"`
	JYConPersonPhone     string      `json:"JYConPersonPhone"`
	JYConPersonFromNum   string      `json:"JYConPersonFromNum"`
	JYConPersonAutoNum   int         `json:"JYConPersonAutoNum"`
	JYConPersonState     string      `json:"JYConPersonState"`
	JYConPersonCode      interface{} `json:"JYConPersonCode"`
	JYConPersonBelongHos string      `json:"JYConPersonBelongHos"`
}

func GetDD(s interface{}) {

	type JsonString struct {
		Code  string        `json:"code"`
		Data  []interface{} `json:"data"`
		Count string        `json:"count"`
		Mesg  string        `json:"mesg"`
	}

}
func JsonStringDateToIntData(c JsonString) JsonInt {
	var depInfo JsonInt
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

func Test() {

}

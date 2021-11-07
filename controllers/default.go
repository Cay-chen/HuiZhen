package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}
type Person struct {
	JYConPersonType      string `json:"JYConPersonType"`
	JYConPersonBelongDep string `json:"JYConPersonBelongDep"`
	Flag                 string `json:"flag"`
	JYConPersonName      string `json:"JYConPersonName"`
	JYConPersonPhone     string `json:"JYConPersonPhone"`
	JYConPersonCode      string `json:"JYConPersonCode"`
	JYConPersonBelongHos string `json:"JYConPersonBelongHos"`
	Mesg                 string `json:"mesg"`
}

type Msg struct {
	Flag string `json:"flag"`
	Mesg string `json:"mesg"`
}

func getMsg(msg string) (person *Msg) {

	res := &Msg{}
	err := json.Unmarshal([]byte(msg), &res)
	if err != nil {
		return
	}
	return res
}

func getPerson(msg string) (person *Person) {

	res := &Person{}
	err := json.Unmarshal([]byte(msg), &res)
	if err != nil {
		return
	}
	return res
}

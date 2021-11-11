package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type ConsultationNoteController struct {
	BaseController
}

func (c *ConsultationNoteController) Get() {
	if c.IsLogin {
		JYConNum := c.GetString("JYConNum")
		if JYConNum != "" {
			info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "", c.PersonUer.JYConPersonCode)
			s := models.FormInfo{}
			err := json.Unmarshal([]byte(info), &s)
			if err != nil {
			}
			c.Data["Form"] = s
			c.TplName = "iframe_consultation_note.html"
		} else {
			c.Redirect("/error/600", 302)
		}
	} else {
		c.Redirect("/error/600", 302)
	}

}
func (c *ConsultationNoteController) Post() {

}

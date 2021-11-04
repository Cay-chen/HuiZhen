package controllers

import (
	"HuiZhen/models"
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
)

type CheckController struct {
	BaseController
}

func (c *CheckController) Get() {
	if c.IsLogin {

		JYConNum := c.GetString("JYConNum")
		if JYConNum != "" {
			info := models.GetFormList("one", "", "", "", "", "", "", JYConNum, "getFormBySomething", "")
			s := models.FormInfo{}
			err := json.Unmarshal([]byte(info), &s)
			if err != nil {
			}
			c.Data["Form"] = s
			JYConFormCreateDate := string([]byte(s.JYConFormCreateDate)[:19])
			c.Data["JYConFormCreateDate"] = JYConFormCreateDate
			JYConDate := string([]byte(s.JYConDate)[:19])
			c.Data["JYConDate"] = JYConDate
			resMsg := models.GetApproveList(JYConNum, "300", "1")
			JYConFormApproveCommentK := ""
			JYConFormApproveCommentY := ""
			JYConFormApproveDateK := ""
			JYConFormApproveDateY := ""
			logs.Debug(resMsg)
			for i := 0; i < len(resMsg.Data); i++ {
				if resMsg.Data[i].JYConFormApproveOldState == "2" {
					JYConFormApproveCommentK = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateK = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
				}
				if resMsg.Data[i].JYConFormApproveOldState == "3" {
					JYConFormApproveCommentY = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateY = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
				}
			}
			c.Data["JYConFormApproveCommentK"] = JYConFormApproveCommentK
			c.Data["JYConFormApproveDateK"] = JYConFormApproveDateK
			c.Data["JYConFormApproveCommentY"] = JYConFormApproveCommentY
			c.Data["JYConFormApproveDateY"] = JYConFormApproveDateY
			c.TplName = "view_page.html"

		} else {
			c.Redirect("/error", 302)
		}

	} else {
		c.Ctx.WriteString("")
	}

}

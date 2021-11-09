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
			JYConFormApproveCommentTY := ""
			JYConFormApproveCommentTK := ""
			JYConFormApproveDateK := ""
			JYConFormApproveDateY := ""
			JYConFormApproveDateTY := ""
			JYConFormApproveDateTK := ""
			logs.Debug(resMsg)
			for i := 0; i < len(resMsg.Data); i++ {
				switch resMsg.Data[i].JYConFormApproveOldState {
				case "2":
					JYConFormApproveCommentK = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateK = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
					break
				case "3":
					JYConFormApproveCommentY = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateY = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
					break
				case "4":
					JYConFormApproveCommentTY = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateTY = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
					break
				case "5":
					JYConFormApproveCommentTK = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateTK = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
					break
				case "6":
					JYConFormApproveCommentY = resMsg.Data[i].JYConFormApproveComment
					JYConFormApproveDateY = string([]byte(resMsg.Data[i].JYConFormApproveDate)[:19])
					break

				}
			}
			if s.JYConType == "2" {

			}
			c.Data["JYConFormApproveCommentK"] = JYConFormApproveCommentK
			c.Data["JYConFormApproveDateK"] = JYConFormApproveDateK
			c.Data["JYConFormApproveCommentY"] = JYConFormApproveCommentY
			c.Data["JYConFormApproveDateY"] = JYConFormApproveDateY
			c.Data["JYConFormApproveCommentTY"] = JYConFormApproveCommentTY
			c.Data["JYConFormApproveDateTY"] = JYConFormApproveDateTY
			c.Data["JYConFormApproveCommentTK"] = JYConFormApproveCommentTK
			c.Data["JYConFormApproveDateTK"] = JYConFormApproveDateTK
			c.Data["ApproveType"] = "平会诊"
			if s.JYConType == "2" {

				c.Data["JYConFormApproveCommentK"] = "/"
				c.Data["JYConFormApproveDateK"] = "/"
				c.Data["JYConFormApproveCommentY"] = "/"
				c.Data["JYConFormApproveDateY"] = "/"
				c.Data["JYConFormApproveCommentTY"] = "/"
				c.Data["JYConFormApproveDateTY"] = "/"
				c.Data["JYConFormApproveCommentTK"] = "/"
				c.Data["JYConFormApproveDateTK"] = "/"
				c.Data["ApproveType"] = "急会诊"
			}
			c.TplName = "check_view_page.html"

		} else {
			c.Redirect("/error", 302)
		}

	} else {
		c.Ctx.WriteString("")
	}

}

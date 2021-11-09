package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type IframeController struct {
	BaseController
}

func (c *IframeController) Get() {
	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "myHz":
			c.Data["A1"] = "{{# if(d.JYConFormPolicy == 1){ }}"
			c.Data["A2"] = "{{# }else{ }}"
			c.Data["A3"] = "{{# } }}"
			c.Data["FormPolicyTpl"] = "    {{#  if(d.JYConFormPolicy == 1){ }}\n    申请状态\n    {{#  } else if(d.JYConFormPolicy == 2){ }}\n    申请方科室审批\n    {{#  } else if(d.JYConFormPolicy == 3){ }}\n    申请方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 4){ }}\n    接收方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 5){ }}\n    接收方科室审批\n    {{#  } else if(d.JYConFormPolicy == 6){ }}\n    审批完成\n    {{# } }}"
			c.Data["FormHos"] = "{{# if(d.JYConOppHos =='WJQY'){ }}\n        温江区医院\n    {{# } else if(d.JYConOppHos =='JYZX'){ }}\n        精医中心\n    {{# } }}"
			c.TplName = "iframe_my_hz.html"
			break
		case "from_dsp":
			if !c.IsYxys {
				IsTO := false
				Type := c.GetString("type")
				c.Data["Type"] = Type
				if c.IsKszr || c.IsFzr {
					if Type == "to" {
						IsTO = true
					}
				}
				docList := models.GetDocList(c.PersonUer.JYConPersonBelongHos, c.PersonUer.JYConPersonBelongDep, c.PersonUer.JYConPersonCode, "200", "1")
				res := models.Person{}
				_ = json.Unmarshal([]byte(docList), &res)
				c.Data["DocList"] = res.Data
				c.Data["IsTO"] = IsTO
				c.Data["FormPolicyTpl"] = "    {{#  if(d.JYConFormPolicy == 1){ }}\n    申请状态\n    {{#  } else if(d.JYConFormPolicy == 2){ }}\n    申请方科室审批\n    {{#  } else if(d.JYConFormPolicy == 3){ }}\n    申请方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 4){ }}\n    接收方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 5){ }}\n    接收方科室审批\n    {{#  } else if(d.JYConFormPolicy == 6){ }}\n    审批完成\n    {{# } }}"
				c.Data["FormHos"] = "{{# if(d.JYConOppHos =='WJQY'){ }}\n        温江区医院\n    {{# } else if(d.JYConOppHos =='JYZX'){ }}\n        精医中心\n    {{# } }}"
				c.TplName = "iframe_d_sp.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "from_ysp":
			if !c.IsYxys {
				Type := c.GetString("type")
				c.Data["approveOldState"] = " {{#  if(d.JYConFormApproveOldState == 1){ }}\n        申请状态\n    {{#  } else if(d.JYConFormApproveOldState == 2){ }}\n        科室审批\n    {{#  } else if(d.JYConFormApproveOldState == 3){ }}\n        申请方医务部审批\n    {{#  } else if(d.JYConFormApproveOldState == 4){ }}\n        接收方医务部审批\n    {{#  } else if(d.JYConFormApproveOldState == 5){ }}\n         接收方科室审批\n    {{# } }}"
				c.Data["approveState"] = "    {{#  if(d.JYConFormApproveState==\"approve\"){ }}\n    同意\n    {{#  } else if(d.JYConFormApproveState==\"reject\"){ }}\n    拒绝\n    {{# } }}\n"
				c.Data["Type"] = Type
				c.TplName = "iframe_y_sp.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "dep_manage":
			if c.IsAdmin || c.IsYwb {
				c.TplName = "iframe_dev_manage.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "person_manage":
			if c.IsAdmin || c.IsYwb {
				c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200").Data
				c.Data["TitleTpl"] = "{{# if(d.JYConPersonIsActive == '已激活') { }}\n    <a  style=\"color: red\">{{d.JYConPersonIsActive}}</a>\n    {{#  } else { }}\n    {{d.JYConPersonIsActive}}\n    {{#  } }}"
				c.TplName = "iframe_person_manage.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "myHzTo":
			c.Data["FormPolicyTpl"] = "    {{#  if(d.JYConFormPolicy == 1){ }}\n    申请状态\n    {{#  } else if(d.JYConFormPolicy == 2){ }}\n    申请方科室审批\n    {{#  } else if(d.JYConFormPolicy == 3){ }}\n    申请方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 4){ }}\n    接收方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 5){ }}\n    接收方科室审批\n    {{#  } else if(d.JYConFormPolicy == 6){ }}\n    审批完成\n    {{# } }}"
			c.Data["FormHos"] = "{{# if(d.JYConOppHos =='WJQY'){ }}\n        温江区医院\n    {{# } else if(d.JYConOppHos =='JYZX'){ }}\n        精医中心\n    {{# } }}"
			c.TplName = "iframe_to_my_hz.html"
			break
		case "extra":
			if c.IsAdmin || c.IsYwb {
				c.Data["PersonType"] = "    {{#  if(d.JYConPersonType == 1){ }}\n        科室主任\n    {{#  } else if(d.JYConPersonType ==2){ }}\n     科室负责人\n    {{#  } else if(d.JYConPersonType == 3){ }}\n     一线医生\n    {{#  } else if(d.JYConPersonType == 4){ }}\n     医务部\n    {{#  } else if(d.JYConPersonType == 0){ }}\n     管理员\n    {{# } }}"
				c.TplName = "iframe_person_extra.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "approve":
			if c.IsAdmin || c.IsYwb {
				JYConNum := c.GetString("JYConNum")
				c.Data["JYConNum"] = JYConNum
				c.Data["approveState"] = "    {{#  if(d.JYConFormApproveState==\"approve\"){ }}\n    同意\n    {{#  } else if(d.JYConFormApproveState==\"reject\"){ }}\n    拒绝\n    {{# } }}\n"
				c.TplName = "iframe_sp_jl.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		default:
			c.Redirect("/error/600", 302)
			break
		}
	} else {
		c.Redirect("/error/600", 302)
	}

}

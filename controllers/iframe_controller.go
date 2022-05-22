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
			/*			c.Data["A1"] = "{{# if(d.JYConFormPolicy == 1){ }}"
						c.Data["A2"] = "{{# }else{ }}"
						c.Data["A3"] = "{{# } }}"*/
			c.Data["BarDemo"] = "    {{# if(d.JYConFormPolicy == 1){ }}\n    <a class=\"layui-btn layui-btn-xs\" lay-event=\"edit\">编辑</a>\n    <a class=\"layui-btn layui-btn-xs\" lay-event=\"submit\">提交</a>\n    <a class=\"layui-btn layui-btn-danger layui-btn-xs\" lay-event=\"del\">删除</a>\n    {{# }else{ }}\n    <a class=\"layui-btn layui-btn-xs layui-btn-disabled\" >编辑</a>\n    <a class=\"layui-btn layui-btn-xs layui-btn-disabled\" >提交</a>\n    <a class=\"layui-btn layui-btn-danger layui-btn-xs layui-btn-disabled\" >删除</a>\n    {{# } }}"
			c.Data["FormPolicyTpl"] = "    {{#  if(d.JYConFormPolicy == 1){ }}\n申请状态\n{{#  } else if(d.JYConFormPolicy == 2){ }}\n申请方科室审批\n{{#  } else if(d.JYConFormPolicy == 3){ }}\n申请方医务部审批\n{{#  } else if(d.JYConFormPolicy == 4){ }}\n接收方医务部审批\n{{#  } else if(d.JYConFormPolicy == 5){ }}\n接收方科室审批\n{{#  } else if(d.JYConFormPolicy == 6){ }}\n待会诊\n{{#  } else if(d.JYConFormPolicy == 7){ }}\n会诊完成\n{{# } }}"
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
				docList := models.GetDocList(c.PersonUer.JYConPersonBelongHos, c.PersonUer.JYConPersonBelongDep, c.PersonUer.JYConPersonCode, "200", "1", c.PersonUer.JYConPersonCode)
				res := models.Person{}
				_ = json.Unmarshal([]byte(docList), &res)
				c.Data["DocList"] = res.Data
				c.Data["IsTO"] = IsTO
				c.Data["FormPolicyTpl"] = "{{#  if(d.JYConFormPolicy == 1){ }}\n申请状态\n{{#  } else if(d.JYConFormPolicy == 2){ }}\n申请方科室审批\n{{#  } else if(d.JYConFormPolicy == 3){ }}\n申请方医务部审批\n{{#  } else if(d.JYConFormPolicy == 4){ }}\n接收方医务部审批\n{{#  } else if(d.JYConFormPolicy == 5){ }}\n接收方科室审批\n{{#  } else if(d.JYConFormPolicy == 6){ }}\n待会诊\n{{#  } else if(d.JYConFormPolicy == 7){ }}\n会诊完成\n{{# } }}"
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
				c.TplName = "iframe_dep_manage.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "person_manage":
			if !c.IsYxys {
				c.Data["TitleTpl"] = "{{# if(d.JYConPersonIsActive == '已激活') { }}\n    <a  style=\"color: red\">{{d.JYConPersonIsActive}}</a>\n    {{#  } else { }}\n    {{d.JYConPersonIsActive}}\n    {{#  } }}"
				if c.IsYwb || c.IsAdmin {
					c.Data["DepList"] = models.GetDepList(c.PersonUer.JYConPersonBelongHos, "1", "200", c.PersonUer.JYConPersonCode).Data
					c.Data["IsYwbAndAdmin"] = true
				} else {
					c.Data["IsYwbAndAdmin"] = false
					res := "[{\"JYConDepCode\": \"" + c.PersonUer.JYConPersonBelongDep + "\", \"JYConDepName\": \"" + c.PersonUer.DepName + "\"}]"
					var res1 []models.Dep
					_ = json.Unmarshal([]byte(res), &res1)
					c.Data["DepList"] = res1
				}
				c.TplName = "iframe_person_manage.html"
			} else {
				c.Redirect("/error/405", 302)
			}
			break
		case "myHzTo":
			if c.IsYxys || c.IsFzr || c.IsKszr {
				c.Data["IsYx"] = true
			} else {
				c.Data["IsYx"] = false
			}
			c.Data["FormPolicyTpl"] = "{{#  if(d.JYConFormPolicy == 1){ }}\n申请状态\n{{#  } else if(d.JYConFormPolicy == 2){ }}\n申请方科室审批\n{{#  } else if(d.JYConFormPolicy == 3){ }}\n申请方医务部审批\n{{#  } else if(d.JYConFormPolicy == 4){ }}\n接收方医务部审批\n{{#  } else if(d.JYConFormPolicy == 5){ }}\n接收方科室审批\n{{#  } else if(d.JYConFormPolicy == 6){ }}\n待会诊\n{{#  } else if(d.JYConFormPolicy == 7){ }}\n会诊完成\n{{# } }}"
			c.Data["FormHos"] = "{{# if(d.JYConOppHos =='WJQY'){ }}\n        温江区医院\n    {{# } else if(d.JYConOppHos =='JYZX'){ }}\n        精医中心\n    {{# } }}"
			c.Data["BarDemo"] = "    {{# if(d.JYConFormPolicy ==6){ }}\n    <a data-method=\"offset\" class=\"layui-btn  layui-btn-xs layui-btn-normal\" data-type=\"auto\" lay-event=\"edit\">写会诊记录</a>\n    {{# }else{ }}\n    <a data-method=\"offset\" class=\"layui-btn  layui-btn-xs layui-btn-disabled\">写会诊记录</a>\n    {{# } }}"
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
			JYConNum := c.GetString("JYConNum")
			c.Data["JYConNum"] = JYConNum
			c.Data["approveState"] = "    {{#  if(d.JYConFormApproveState==\"approve\"){ }}\n    同意\n    {{#  } else if(d.JYConFormApproveState==\"reject\"){ }}\n    拒绝\n    {{# } }}\n"
			c.TplName = "iframe_sp_jl.html"
			break
		default:
			c.Redirect("/error/600", 302)
			break
		}
	} else {
		c.Redirect("/error/600", 302)
	}

}

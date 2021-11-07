package controllers

import (
	"HuiZhen/models"
	"encoding/json"
)

type TableController struct {
	BaseController
}

func (c *TableController) Get() {
	if c.IsLogin {

		c.Data["A1"] = "{{# if(d.JYConFormPolicy == 1){ }}"
		c.Data["A2"] = "{{# }else{ }}"
		c.Data["A3"] = "{{# } }}"
		c.Data["A4"] = " {{#  if(d.JYConFormPolicy == 1){ }}\n        申请状态\n    {{#  } else if(d.JYConFormPolicy == 2){ }}\n        申请方科室审批\n    {{#  } else if(d.JYConFormPolicy == 3){ }}\n        申请方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 4){ }}\n        接收方医务部审批\n    {{#  } else if(d.JYConFormPolicy == 5){ }}\n         接收方科室审批\n    {{# } }}"
		//c.Data["A5"]="{{#  if(d.JYConOppHos == 'WJQY'){ }}\n    温江区医院\n    {{#  } else if(d.JYConFormPolicy == 'JYZX'){ }}\n    精医中心\n    {{# } }}"
		c.Data["A5"] = "{{#  if(d.JYConOppHos == "
		c.Data["A6"] = "){ }}\n    温江区医院\n    {{#  } else if(d.JYConOppHos == "
		c.Data["A7"] = "){ }}\n    精医中心\n    {{# } }}"

		c.Data["approveOldState"] = " {{#  if(d.JYConFormApproveOldState == 1){ }}\n        申请状态\n    {{#  } else if(d.JYConFormApproveOldState == 2){ }}\n        科室审批\n    {{#  } else if(d.JYConFormApproveOldState == 3){ }}\n        申请方医务部审批\n    {{#  } else if(d.JYConFormApproveOldState == 4){ }}\n        接收方医务部审批\n    {{#  } else if(d.JYConFormApproveOldState == 5){ }}\n         接收方科室审批\n    {{# } }}"
		c.Data["approveState"] = "    {{#  if(d.JYConFormApproveState==\"approve\"){ }}\n    通过\n    {{#  } else if(d.JYConFormApproveState==\"reject\"){ }}\n    申请方科室审批\n    {{# } }}\n"
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "my":
			c.TplName = "table.html"
			break
		case "from_dsp":
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
			c.TplName = "table_from.html"
			break
		case "from_ysp":
			Type := c.GetString("type")
			c.Data["Type"] = Type
			c.TplName = "table_approve.html"
			break
		}
	} else {
		c.Redirect("/error/1245", 302)

	}

}

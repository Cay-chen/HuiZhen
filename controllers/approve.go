package controllers

type ApproveController struct {
	BaseController
}

func (c *ApproveController) Get() {
	if c.IsLogin {
		c.TplName = "approve.html"

	} else {
		c.TplName = "error.html"
	}

}

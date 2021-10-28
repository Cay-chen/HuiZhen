package controllers

type PageController struct {
	BaseController
}

func (c *PageController) Get() {
	if c.IsLogin {

	} else {
		c.Ctx.Redirect(302, "/error")
	}

}

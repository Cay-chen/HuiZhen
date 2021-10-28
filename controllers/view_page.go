package controllers

type ViewPageController struct {
	BaseController
}

func (c *ViewPageController) Get() {
	c.TplName = "view_page.html"
}
